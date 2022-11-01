package server

import (
	"testing"
	"time"
)

func TestConfigServer(t *testing.T) {
	server := NewConfigServer(":3333", []RequestHandler{NewParameterHandler("/", nil)}, nil)
	go server.ListenAndServe()
	server.Shutdown()
}

func TestConfigServerFail(t *testing.T) {
	server := NewConfigServer(":3333", []RequestHandler{}, nil)
	defer func() {
		_ = recover()
		server.Shutdown()
	}()
	server2 := NewConfigServer(":3333", []RequestHandler{}, nil)
	go server.ListenAndServe()
	time.Sleep(100 * time.Millisecond)
	server2.ListenAndServe()
	t.Fatal()
}
