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
	log.Printf("got request to: %v\n", request.RequestURI)
	writer.Header().Set("Content-Type", "application/json")
	result, err := handler.Process(request)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	_, err = writer.Write(result)
	if err != nil {
		log.Printf("could not respond to: %v, error: %v", request.RemoteAddr, err)
		return
	}
	log.Printf("responded to: %v with: %v\n", request.RemoteAddr, string(result))
}

func (handler *DefaultRequestHandler) Path() string {
	return handler.path
}
