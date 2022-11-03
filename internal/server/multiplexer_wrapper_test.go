package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountingMultiplexerHandle(t *testing.T) {
	multiplexer := CountingMultiplexer{
		Multiplexer:   http.NewServeMux(),
		totalHandlers: 0,
	}
	multiplexer.Handle("/", NewParameterHandler("", nil))
	if multiplexer.totalHandlers == 0 {
		t.Fatal()
	}
}

func TestSafeMultiplexerHandle(t *testing.T) {
	multiplexer := SafeMultiplexer{Multiplexer: http.NewServeMux()}
	multiplexer.Handle("/", NewParameterHandler("/", nil))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	multiplexer.ServeHTTP(w, req) // should not fail
}

func TestWrappingMultiplexers(t *testing.T) {
	safeMultiplexer := &SafeMultiplexer{Multiplexer: http.NewServeMux()}
	countingMultiplexer := CountingMultiplexer{
		Multiplexer:   safeMultiplexer,
		totalHandlers: 0,
	}

	countingMultiplexer.Handle("/", NewParameterHandler("/", nil))
	if countingMultiplexer.totalHandlers == 0 {
		t.Fatal()
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	countingMultiplexer.ServeHTTP(w, req) // should not fail
}

func TestNewSafeCountingMultiplexer(t *testing.T) {
	multiplexer := NewSafeCountingMultiplexer()

	multiplexer.Handle("/", NewParameterHandler("/", nil))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	multiplexer.ServeHTTP(w, req) // should not fail
}
