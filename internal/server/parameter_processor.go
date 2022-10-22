package server

import (
	"config_master/internal/parameters"
	"encoding/json"
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
			val := map[string]string{"error": err.Error()}
			res, _ := json.Marshal(val)
			return res
		}
		err = handler.Set(data)
		if err != nil {
			val := map[string]string{"error": err.Error()}
			res, _ := json.Marshal(val)
			return res
		}
		return []byte(`{"result":"OK"}`)
	default:
		return handler.GetAsJSON()
	}
}

func NewParameterHandler(path string, parameter parameters.Parameter) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &ParameterProcessor{parameter}}
}
