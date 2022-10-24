package server

import (
	"config_master/internal/parameters"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ParameterProcessor struct {
	parameters.Parameter
}

func (handler *ParameterProcessor) Process(request *http.Request) []byte {
	switch request.Method {
	case http.MethodPut:
		data, err := io.ReadAll(request.Body)
		if err != nil {
			return parseError(err)
		}
		err = handler.Set(data)
		if err != nil {
			return parseError(err)
		}
		return []byte(`{"result":"OK"}`)
	case http.MethodGet:
		return handler.GetAsJSON()
	}
	return parseError(fmt.Errorf("method %v not supported", request.Method))
}

func NewParameterHandler(path string, parameter parameters.Parameter) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &ParameterProcessor{parameter}}
}

func parseError(err error) []byte {
	val := map[string]string{"error": err.Error()}
	res, _ := json.Marshal(val)
	return res
}
