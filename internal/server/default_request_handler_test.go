package server

import (
	"config_master/internal/parameters"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultRequestHandler(t *testing.T) {
	handler := NewParameterHandler("/", &ParameterProcessor{Parameter: parameters.NewSimpleParameter("test", 1)})
	var expected = 200
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != expected {
		t.Fatalf("expected status code to be %v, got %v", expected, res.StatusCode)
	}
}
