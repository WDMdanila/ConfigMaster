package server

import "testing"

func TestConfigServer(t *testing.T) {
	server := NewConfigServer(":3333", []RequestHandler{})
	go server.ListenAndServe()
	server.Shutdown()
}
