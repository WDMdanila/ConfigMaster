package main

import (
	"config_master/internal/parameters"
	"config_master/internal/server"
	"config_master/internal/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
)

func extractConfigFileNameAndPath(fileName string) (string, string) {
	fileName = strings.ReplaceAll(fileName, `\`, "/")
	fileName = strings.ReplaceAll(fileName, "//", "/")
	log.Printf("found config: %v\n", fileName)
	folderNameIndex := strings.Index(fileName, "/")
	return fileName, utils.GetFilenameWithoutExt(fileName[folderNameIndex+1:])
}

func createConfigHandlers(configDirectory string, strictTypes bool) []server.RequestHandler {
	var handlers []server.RequestHandler
	for _, configFile := range utils.FindFilesWithExtInDirectory(configDirectory, "json") {
		var configHttpPath string
		configFile, configHttpPath = extractConfigFileNameAndPath(configFile)
		paramReader := parameters.NewJSONParameterReader(configFile, strictTypes)
		parametersMap := paramReader.Read()
		for key, value := range parametersMap {
			handlerPath := fmt.Sprintf("/%v/%v", configHttpPath, key)
			handlers = append(handlers, server.NewParameterHandler(handlerPath, value))
		}
	}
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
	handlers := createConfigHandlers(configDir, strictTypes)
	configServer := server.NewConfigServer(address, handlers)
	defer configServer.Shutdown()
	go configServer.ListenAndServe()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}
