package main

import (
	"config_master/internal/parameters"
	"config_master/internal/server"
	"config_master/internal/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
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

func main() {
	handlers := createConfigHandlers("./configs")
	handlers = append(handlers, createTimestampHandler())
	configServer := server.NewConfigServer(":3333", handlers)
	defer configServer.Shutdown()
	go configServer.ListenAndServe()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}
