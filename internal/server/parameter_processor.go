package server

import (
	"config_master/internal/parameters"
	"net/http"
)

type ParameterProcessor struct {
	parameters.Parameter
}

func (handler *ParameterProcessor) Process(*http.Request) []byte {
	return handler.GetAsJSON()
}

func NewParameterHandler(path string, parameter parameters.Parameter) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &ParameterProcessor{parameter}}
}
