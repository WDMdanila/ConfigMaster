package server

import (
	"config_master/internal/parameters"
	"config_master/internal/utils"
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
		return handler.handlePUT(request)
	case http.MethodGet:
		return parseResponse("value", handler.Value())
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func (handler *ParameterProcessor) handlePUT(request *http.Request) []byte {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	value, err := utils.ExtractFromJSON[interface{}](data, "value")
	if err != nil {
		return parseResponse("error", err.Error())
	}
	err = handler.Set(value)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	return parseResponse("result", "OK")
}

func parseResponse(name string, value interface{}) []byte {
	val := map[string]interface{}{name: value}
	res, _ := json.Marshal(val)
	return res
}

func NewParameterHandler(path string, parameter parameters.Parameter) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &ParameterProcessor{parameter}}
}
