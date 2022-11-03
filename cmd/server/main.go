package main

import (
	"config_master/internal/parameters"
	"config_master/internal/server"
	"config_master/internal/utils"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func createRequestHandlers(configDirectory string, strictTypes bool, multiplexer server.Multiplexer) []server.RequestHandler {
	var handlers []server.RequestHandler
	var nestedHandlers []server.RequestHandler
	configFiles, err := utils.FindFilesWithExtRecursively(configDirectory, "json")
	if err != nil {
		panic(err)
	}
	for _, configFile := range configFiles {
		var configHttpPath string
		processors := make([]server.RequestHandler, 0)
		configFile, configHttpPath = utils.ExtractFileNameAndPath(configFile)
		paramReader := parameters.NewJSONParameterReader(configFile, strictTypes)
		parametersMap := paramReader.Read()
		for key, value := range parametersMap {
			handlerPath := fmt.Sprintf("/%v/%v", configHttpPath, key)
			handler := server.NewParameterHandler(handlerPath, value)
			processors = append(processors, handler)
			handlers = append(handlers, handler)
		}
		nestedHandler := server.NewNestedRequestHandler("/"+configHttpPath, processors, multiplexer)
		handlers = append(handlers, nestedHandler)
		nestedHandlers = append(nestedHandlers, nestedHandler)
	}
	handlers = append(handlers, server.NewNestedRequestHandler("/", nestedHandlers, multiplexer))
	return handlers
}

func parseArgs() (string, string, bool) {
	address := flag.String("address", "", "address to use")
	port := flag.Int64("port", 3333, "port to use")
	configDir := flag.String("config-dir", "./configs", "path to directory with configs")
	strictTypes := flag.Bool("strict", false, "disallow parameters' type changes")
	flag.Parse()
	finalAddress := fmt.Sprintf("%v:%v", *address, *port)
	return finalAddress, *configDir, *strictTypes
}

func main() {
	address, configDir, strictTypes := parseArgs()
	multiplexer := server.NewSafeCountingMultiplexer()
	handlers := createRequestHandlers(configDir, strictTypes, multiplexer)
	configServer := server.NewConfigServer(address, handlers, multiplexer)
	defer configServer.Shutdown()
	go configServer.ListenAndServe()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}
