package server

import (
	"bytes"
	"config_master/internal/parameters"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewNestedRequestHandler(t *testing.T) {
	handler1 := NewNestedRequestHandler("/handler_1", nil)
	handler2 := NewParameterHandler("/handler_1/handler_2", parameters.NewSimpleParameter("param_name", 1))
	handler1.AddProcessor(handler2)
	res := handler1.Process(nil)
	if len(res) != 1 {
		t.Fatal()
	}
	out, _ := json.Marshal(res)
	fmt.Printf("%v\n", string(out))
}

func TestNestedRequestHandlerServeHTTP(t *testing.T) {
	expected := []byte(`{"/handler_1":{"param_name":1}}`)
	req := httptest.NewRequest(http.MethodGet, "/handler_1", nil)
	w := httptest.NewRecorder()
	handler := NewNestedRequestHandler("/handler_1", nil)
	handler2 := NewParameterHandler("/handler_1/handler_2", parameters.NewSimpleParameter("param_name", 1))
	handler.AddProcessor(handler2)
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

func TestNestedRequestHandlerServeHTTP404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/qwe", nil)
	w := httptest.NewRecorder()
	handler := NewNestedRequestHandler("/", nil)
	handler2 := NewParameterHandler("/handler_2", parameters.NewSimpleParameter("param_name", 1))
	handler.AddProcessor(handler2)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatal("Expected 404 not found")
	}
}
