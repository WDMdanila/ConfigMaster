package server

import (
	"context"
	"log"
	"net/http"
)

type ConfigServer struct {
	http.Server
}

func (s *ConfigServer) Shutdown() {
	log.Printf("closing server gracefuly\n")
	err := s.Server.Shutdown(context.Background())
	if err != nil {
		log.Panicf("could not shutdown gracefuly, error: %v", err)
	}
}

func (s *ConfigServer) ListenAndServe() {
	log.Printf("Listening on %v\n", s.Addr)
	err := s.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Printf("closed server gracefuly\n")
	} else if err != nil {
		log.Panicf("could not start server, error: %v\n", err)
	}
}

func NewConfigServer(address string, handlers []RequestHandler) *ConfigServer {
	m := http.NewServeMux()
	configServer := &ConfigServer{http.Server{Addr: address, Handler: m}}
	for index, handler := range handlers {
		m.Handle(handler.Path(), handler)
		log.Printf("registered %v handler under: %v", index+1, handler.Path())
	}
	return configServer
}
