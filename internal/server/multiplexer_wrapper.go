package server

import (
	"log"
	"net/http"
)

type Multiplexer interface {
	Handle(string, http.Handler)
	http.Handler
}

type SafeMultiplexer struct {
	Multiplexer
}

func (m *SafeMultiplexer) Handle(path string, handler http.Handler) {
	m.Multiplexer.Handle(path, NewHandlerWrapper(path, handler))
}

type CountingMultiplexer struct {
	Multiplexer
	totalHandlers int
}

func (m *CountingMultiplexer) Handle(path string, handler http.Handler) {
	m.Multiplexer.Handle(path, handler)
	m.totalHandlers++
	log.Printf("registered %v handler under: %v", m.totalHandlers, path)
}

func NewSafeCountingMultiplexer() Multiplexer {
	return &SafeMultiplexer{Multiplexer: &CountingMultiplexer{Multiplexer: http.NewServeMux(), totalHandlers: 0}}
}
