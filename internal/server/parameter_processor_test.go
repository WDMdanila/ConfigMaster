package server

import (
	"bytes"
	"config_master/internal/parameters"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParameterHandler(t *testing.T) {
	expected := []byte(`{"value":1}`)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler := NewParameterHandler("/", parameters.NewSimpleParameter("value", 1))
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

func TestParameterHandlerFail(t *testing.T) {
	expected := []byte(`{"error":"method POST not supported"}`)
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	handler := NewParameterHandler("/", parameters.NewSimpleParameter("value", 1))
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

func TestParameterHandlerPut(t *testing.T) {
	expected := []byte(`{"result":"OK"}`)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(`{"value": 1}`)))
	w := httptest.NewRecorder()
	handler := NewParameterHandler("/", parameters.NewSimpleParameter("value", 1))
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	if !bytes.Equal(data, expected) {
		t.Fatalf(`expected %v got %v`, string(expected), string(data))
	}
}

func TestParameterHandlerPutFail(t *testing.T) {
	defer func() { _ = recover() }()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()
	handler := NewParameterHandler("/", parameters.NewSimpleParameter("value", 1))
	handler.ServeHTTP(w, req)
}

func TestParameterHandlerPutFail2(t *testing.T) {
	expected := []byte(`{"error":"failed to set value, error: could not parse 1, got type float64 but string was expected"}`)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(`{"value": 1}`)))
	w := httptest.NewRecorder()
	handler := NewParameterHandler("/", parameters.NewSimpleStrictParameter("value", "1"))
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(resp, expected) {
		t.Fatalf(`expected %v got %v`, string(expected), string(resp))
	}
}
