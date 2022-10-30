package server

import (
	"log"
	"net/http"
)

type DefaultRequestHandler struct {
	ReachableRequestHandler
	Processor
}

func (h *DefaultRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	writer.Header().Set("Content-Type", "application/json")
	result := h.Process(request)
	_, err := writer.Write(result)
	if err != nil {
		log.Printf("could not respond to: %v, error: %v", request.RemoteAddr, err)
		return
	}
	log.Printf("responded to: %v with: %v\n", request.RemoteAddr, string(result))
}
