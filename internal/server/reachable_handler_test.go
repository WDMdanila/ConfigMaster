package server

import "testing"

func TestReachableRequestHandlerPath(t *testing.T) {
	handler := ReachableRequestHandler{path: "/"}
	if handler.Path() != "/" {
		t.Fatalf("Expected: %v, got: %v", "/", handler.Path())
	}
}
