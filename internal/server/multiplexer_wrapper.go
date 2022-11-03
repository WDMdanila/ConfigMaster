package server

import (
	"log"
	"net/http"
)

type Multiplexer interface {
	Handle(string, http.Handler)
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
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

func (m *CountingMultiplexer) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	m.Multiplexer.HandleFunc(path, handler)
	m.totalHandlers++
	log.Printf("registered %v handler under: %v", m.totalHandlers, path)
}

func NewSafeCountingMultiplexer() Multiplexer {
	return &SafeMultiplexer{Multiplexer: &CountingMultiplexer{Multiplexer: http.NewServeMux(), totalHandlers: 0}}
}
