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
	"path/filepath"
	"strconv"
)

func createConfigHandlers(configDirectory string) []server.RequestHandler {
	var handlers []server.RequestHandler
	for _, configFile := range utils.FindFilesWithExtInDirectory(configDirectory, "json") {
		configPath := filepath.Join(configDirectory, configFile)
		log.Printf("found config: %v\n", configPath)
		paramReader := parameters.NewJSONParameterReader(configPath)
		parametersMap := paramReader.Read()
		for key, value := range parametersMap {
			handlerPath := fmt.Sprintf("/%v/%v", utils.GetFilenameWithoutExt(configFile), key)
			handlers = append(handlers, server.NewParameterHandler(handlerPath, value))
		}
	}
	return handlers
}

func createTimestampHandler() server.RequestHandler {
	return server.NewTimestampHandler("/timestamp", 1_000_000_000)
}

func parseArgs() (string, string) {
	address := flag.String("address", "", "address to use")
	port := flag.Int64("port", 3333, "port to use")
	configDir := flag.String("config-dir", "./configs", "path to directory with configs")
	flag.Parse()
	finalAddress := *address + ":" + strconv.FormatInt(*port, 10)
	return finalAddress, *configDir
}

func main() {
	address, configDir := parseArgs()
	handlers := createConfigHandlers(configDir)
	handlers = append(handlers, createTimestampHandler())
	configServer := server.NewConfigServer(address, handlers)
	defer configServer.Shutdown()
	go configServer.ListenAndServe()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}
