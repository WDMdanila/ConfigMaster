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

var requestMethodHandlerFunc = map[string]func(*http.Request, *ParameterProcessor) []byte{
	"":             handleGET,
	http.MethodGet: handleGET,
	http.MethodPut: handlePUT,
}

func (processor *ParameterProcessor) Process(request *http.Request) []byte {
	if val, ok := requestMethodHandlerFunc[request.Method]; ok {
		return val(request, processor)
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func handleGET(_ *http.Request, processor *ParameterProcessor) []byte {
	return parseResponse("value", processor.Value())
}

func handlePUT(request *http.Request, processor *ParameterProcessor) []byte {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	value, err := utils.ExtractFromJSON[interface{}](data, "value")
	if err != nil {
		return parseResponse("error", err.Error())
	}
	err = processor.Set(value)
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
	return &DefaultRequestHandler{ReachableRequestHandler: ReachableRequestHandler{path: path}, Processor: &ParameterProcessor{parameter}}
}
