package server

import (
	"config_master/internal/parameters"
	"encoding/json"
	"fmt"
	"net/http"
)

type ParameterHandler struct {
	ReachableRequestHandler
	parameters.Parameter
}

func (p *ParameterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	result := p.GetResponse(request)
	_, _ = writer.Write(result)
}

var parameterHandlerFunctions = map[string]HandlerFunction[ParameterHandler]{
	"":              handleGET,
	http.MethodGet:  handleGET,
	http.MethodPut:  handlePUT,
	http.MethodPost: handlePOST,
}

func (p *ParameterHandler) GetResponse(request *http.Request) []byte {
	if val, ok := parameterHandlerFunctions[request.Method]; ok {
		return val(p, request)
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func handleGET(processor *ParameterHandler, _ *http.Request) []byte {
	res, err := json.Marshal(processor.Value())
	if err != nil {
		return parseResponse("error", err)
	}
	return res
}

func handlePUT(processor *ParameterHandler, request *http.Request) []byte {
	value, err := extractData(request)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	err = processor.Set(value["value"])
	if err != nil {
		return parseResponse("error", err.Error())
	}
	return parseResponse("result", "OK")
}

func handlePOST(processor *ParameterHandler, request *http.Request) []byte {
	value, err := extractData(request)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	switch parameter := processor.Parameter.(type) {
	case parameters.Extender:
		parameter.Extend(value["value"])
	default:
		return parseResponse("error", fmt.Sprintf("parameter %v is not a selection parameter", parameter.Name()))
	}
	return parseResponse("result", "OK")
}

func NewParameterHandler(path string, parameter parameters.Parameter) *ParameterHandler {
	return &ParameterHandler{
		ReachableRequestHandler: ReachableRequestHandler{path: path},
		Parameter:               parameter,
	}
}
