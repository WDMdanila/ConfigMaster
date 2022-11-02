package server

import (
	"log"
	"net/http"
)

type RecoveryHandler struct {
	handler RequestHandler
}

func (h *RecoveryHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	defer handleError(writer, request)()
	if request.URL.Path != h.handler.Path() {
		log.Printf("responded to: %v with: 404 not found\n", request.RemoteAddr)
		http.NotFound(writer, request)
		return
	}
	h.handler.ServeHTTP(writer, request)
}

func handleError(writer http.ResponseWriter, request *http.Request) func() {
	return func() {
		if err := recover(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			data := parseResponse("error", err)
			_, err = writer.Write(data)
			if err != nil {
				log.Printf("could not respond to: %v, error: %v", request.RemoteAddr, err)
				return
			}
			log.Printf("responded to: %v with: %v\n", request.RemoteAddr, string(data))
		}
	}
}

func NewRecoveryHandler(handler RequestHandler) *RecoveryHandler {
	return &RecoveryHandler{handler: handler}
}
