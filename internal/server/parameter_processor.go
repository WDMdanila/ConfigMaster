package server

import (
	"config_master/internal/parameters"
	"io"
	"net/http"
)

type ParameterProcessor struct {
	parameters.Parameter
}

func (handler *ParameterProcessor) Process(request *http.Request) []byte {
	switch request.Method {
	case "POST":
		data, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		handler.Set(data)
		return []byte(`{"result": "OK"}`)
	default:
		return handler.GetAsJSON()
	}
}

func NewParameterHandler(path string, parameter parameters.Parameter) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &ParameterProcessor{parameter}}
}
