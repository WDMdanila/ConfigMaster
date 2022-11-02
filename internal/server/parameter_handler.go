package server

import (
	"config_master/internal/parameters"
	"config_master/internal/utils"
	"fmt"
	"io"
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

var requestMethodHandlerFunc = map[string]func(*http.Request, *ParameterHandler) []byte{
	"":             handleGET,
	http.MethodGet: handleGET,
	http.MethodPut: handlePUT,
}

func (p *ParameterHandler) GetResponse(request *http.Request) []byte {
	if val, ok := requestMethodHandlerFunc[request.Method]; ok {
		return val(request, p)
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func handleGET(_ *http.Request, processor *ParameterHandler) []byte {
	return parseResponse("value", processor.Value())
}

func handlePUT(request *http.Request, processor *ParameterHandler) []byte {
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
	res, _ := utils.GetAsJSON(name, value)
	return res
}

func NewParameterHandler(path string, parameter parameters.Parameter) *ParameterHandler {
	return &ParameterHandler{
		ReachableRequestHandler: ReachableRequestHandler{path: path},
		Parameter:               parameter,
	}
}
