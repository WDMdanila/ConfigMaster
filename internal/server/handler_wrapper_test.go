package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryHandlerShouldRecover(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/handler", nil)
	w := httptest.NewRecorder()
	handler := NewHandlerWrapper("", nil)
	handler.ServeHTTP(w, req)
}

func TestRecoveryHandlerShouldRecoverIfErrorInHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/handler", nil)
	w := httptest.NewRecorder()
	handler := NewHandlerWrapper("/handler", NewParameterHandler("/handler", nil))
	handler.ServeHTTP(w, req)
}

func TestRecoveryHandlerShouldReturn404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/handler", nil)
	w := httptest.NewRecorder()
	handler := NewHandlerWrapper("/", NewParameterHandler("/", nil))
	handler.ServeHTTP(w, req)
	if w.Result().StatusCode != 404 {
		t.Fatal()
	}
}
