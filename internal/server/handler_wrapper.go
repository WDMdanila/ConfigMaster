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
	ReachableRequestHandler
}

func (h *HandlerWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	loggingWriter := NewLoggingResponseWriter(writer, request.RemoteAddr)
	defer handleError(loggingWriter)
	if request.URL.Path != h.Path() {
		http.NotFound(loggingWriter, request)
		return
	}
	h.Handler.ServeHTTP(loggingWriter, request)
}

func handleError(writer http.ResponseWriter) {
	if err := recover(); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		data := parseResponse("error", err)
		_, _ = writer.Write(data)
	}
}

func NewHandlerWrapper(path string, handler http.Handler) *HandlerWrapper {
	return &HandlerWrapper{ReachableRequestHandler: ReachableRequestHandler{path: path, Handler: handler}}
}

func WrapHandler(handler RequestHandler) *HandlerWrapper {
	return NewHandlerWrapper(handler.Path(), handler)
}

func NewLoggingResponseWriter(writer http.ResponseWriter, remoteAddr string) *LoggingResponseWriter {
	return &LoggingResponseWriter{ResponseWriter: writer, RemoteAddr: remoteAddr}
}
