package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimestampHandler(t *testing.T) {
	expected := []byte(`{"value":1010}`)
	req := httptest.NewRequest(http.MethodGet, "/timestamp?timestamp=1000", nil)
	w := httptest.NewRecorder()
	handler := NewTimestampHandler("/", 10)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	if !bytes.Equal(data, expected) {
		t.Fatalf(`expected {"value":1} got %v`, string(data))
	}
}

func TestTimestampHandlerMultiple(t *testing.T) {
	expected := []byte(`{"value":1010}`)
	req := httptest.NewRequest(http.MethodGet, "/timestamp?timestamp=1000", nil)
	w := httptest.NewRecorder()
	handler := NewTimestampHandler("/", 10)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	req = httptest.NewRequest(http.MethodGet, "/timestamp?timestamp=10", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	if !bytes.Equal(data, expected) {
		t.Fatalf(`expected {"value":1} got %v`, string(data))
	}
}

func TestTimestampHandlerShouldReturn500(t *testing.T) {
	expected := 500
	req := httptest.NewRequest(http.MethodGet, "/timestamp", nil)
	w := httptest.NewRecorder()
	handler := NewTimestampHandler("/", 10)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != expected {
		t.Fatalf("expected status code to be %v, got %v", expected, res.StatusCode)
	}
}
