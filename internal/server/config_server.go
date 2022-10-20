package server

import (
	"context"
	"log"
	"net/http"
)

type ConfigServer struct {
	http.Server
}

func (configServer *ConfigServer) Shutdown() {
	log.Printf("closing server gracefuly\n")
	err := configServer.Server.Shutdown(context.Background())
	if err != nil {
		log.Panicf("could not shutdown gracefuly, error: %v", err)
	}
}

func (configServer *ConfigServer) ListenAndServe() {
	log.Printf("Listening on %v\n", configServer.Addr)
	err := configServer.Server.ListenAndServe()
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
