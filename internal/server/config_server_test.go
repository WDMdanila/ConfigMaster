package server

import "testing"

func TestConfigServer(t *testing.T) {
	server := NewConfigServer(":3333", []RequestHandler{&DefaultRequestHandler{path: "/"}})
	go server.ListenAndServe()
	server.Shutdown()
}
