package server

import (
	"log"
	"net/http"
	"strings"
)

type LoggingResponseWriter struct {
	RemoteAddr string
	http.ResponseWriter
}

func (w *LoggingResponseWriter) Write(data []byte) (int, error) {
	wroteNum, err := w.ResponseWriter.Write(data)
	if err != nil {
		log.Printf("could not respond to: %v, error: %v", w.RemoteAddr, err)
		return wroteNum, err
	}
	log.Printf("responded to: %v with: %v\n", w.RemoteAddr, strings.TrimSuffix(string(data), "\n"))
	return wroteNum, nil
}

type HandlerWrapper struct {
	handler RequestHandler
}

func (h *HandlerWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	loggingWriter := &LoggingResponseWriter{ResponseWriter: writer, RemoteAddr: request.RemoteAddr}
	defer handleError(loggingWriter)()
	if request.URL.Path != h.handler.Path() {
		http.NotFound(loggingWriter, request)
		return
	}
	h.handler.ServeHTTP(loggingWriter, request)
}

func handleError(writer http.ResponseWriter) func() {
	return func() {
		if err := recover(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			data := parseResponse("error", err)
			_, _ = writer.Write(data)
		}
	}
}

func NewRecoveryHandler(handler RequestHandler) *HandlerWrapper {
	return &HandlerWrapper{handler: handler}
}
