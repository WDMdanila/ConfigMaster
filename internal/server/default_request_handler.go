package server

import (
	"log"
	"net/http"
)

type DefaultRequestHandler struct {
	path string
	Processor
}

func (handler *DefaultRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	writer.Header().Set("Content-Type", "application/json")
	result := handler.Process(request)
	_, err := writer.Write(result)
	if err != nil {
		log.Printf("could not respond to: %v, error: %v", request.RemoteAddr, err)
		return
	}
	log.Printf("responded to: %v with: %v\n", request.RemoteAddr, string(result))
}

func (handler *DefaultRequestHandler) Path() string {
	return handler.path
}
