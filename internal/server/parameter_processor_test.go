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
	handler := NewParameterHandler("/", &parameters.SimpleParameter[int]{Value: 1})
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
