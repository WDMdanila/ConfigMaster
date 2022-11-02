package server

import (
	"config_master/internal/parameters"
	"config_master/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NestedRequestHandler struct {
	ReachableRequestHandler
	processors  []RequestHandler
	multiplexer *http.ServeMux
}

func (h *NestedRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := h.GetResponse(request)
	_, _ = writer.Write(data)
}

func (h *NestedRequestHandler) GetResponse(request *http.Request) []byte {
	switch request.Method {
	case http.MethodGet:
		result := h.Process(request)
		data, _ := json.Marshal(result)
		return data
	case http.MethodPut:
		data, err := io.ReadAll(request.Body)
		if err != nil {
			return parseResponse("error", err.Error())
		}
		value, err := utils.DecodeJSON[map[string]interface{}](data)
		if err != nil {
			return parseResponse("error", err.Error())
		}
		for key, val := range value {
			switch v := val.(type) {
			case map[string]interface{}:
				parameter := parameters.FromJSON(key, v, false)
				paramHandler := NewParameterHandler(h.Path()+"/"+key, parameter)
				h.processors = append(h.processors, paramHandler)
				h.multiplexer.Handle(paramHandler.Path(), paramHandler)
				log.Printf("registered parameter %v on %v", key, paramHandler.Path())
			default:
				parameter := parameters.FromJSON(key, value, false)
				paramHandler := NewParameterHandler(h.Path()+"/"+key, parameter)
				h.processors = append(h.processors, paramHandler)
				h.multiplexer.Handle(paramHandler.Path(), paramHandler)
				log.Printf("registered parameter %v on %v", key, paramHandler.Path())
			}

		}
		return parseResponse("result", "OK")
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func (h *NestedRequestHandler) Process(request *http.Request) map[string]interface{} {
	tmp := map[string]interface{}{}
	for _, processor := range h.processors {
		for key, val := range processor.Process(request) {
			tmp[key] = val
		}
	}
	res := map[string]interface{}{h.Path(): tmp}
	return res
}

func (h *NestedRequestHandler) AddProcessor(processor RequestHandler) {
	h.processors = append(h.processors, processor)
}

func NewNestedRequestHandler(path string, processors []RequestHandler, multiplexer *http.ServeMux) *NestedRequestHandler {
	return &NestedRequestHandler{ReachableRequestHandler: ReachableRequestHandler{path: path}, processors: processors, multiplexer: multiplexer}
}
