package server

import (
	"config_master/internal/parameters"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NestedRequestHandler struct {
	ReachableRequestHandler
	handlers    []RequestHandler
	multiplexer Multiplexer
}

var nestedHandlerFunctions = map[string]func(*NestedRequestHandler, *http.Request) []byte{
	"":              nestedProcessGet,
	http.MethodGet:  nestedProcessGet,
	http.MethodPost: nestedProcessPost,
}

func nestedProcessGet(h *NestedRequestHandler, _ *http.Request) []byte {
	result := h.Describe()
	data, _ := json.Marshal(result)
	return data
}

func nestedProcessPost(h *NestedRequestHandler, request *http.Request) []byte {
	value, err := extractData(request)
	if err != nil {
		return parseResponse("error", err.Error())
	}
	for key, val := range value {
		switch v := val.(type) {
		case map[string]interface{}:
			parameter := parameters.FromJSON(key, v, false)
			h.registerHandler(key, parameter)
		default:
			parameter := parameters.NewSimpleParameter(key, v)
			h.registerHandler(key, parameter)
		}
	}
	return parseResponse("result", "OK")
}

func (h *NestedRequestHandler) registerHandler(key string, parameter parameters.Parameter) {
	paramHandler := NewParameterHandler(h.Path()+"/"+key, parameter)
	h.AddProcessor(paramHandler)
	log.Printf("registered parameter %v on %v", key, paramHandler.Path())
}

func (h *NestedRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := h.GetResponse(request)
	_, _ = writer.Write(data)
}

func (h *NestedRequestHandler) GetResponse(request *http.Request) []byte {
	if val, ok := nestedHandlerFunctions[request.Method]; ok {
		return val(h, request)
	}
	return parseResponse("error", fmt.Sprintf("method %v not supported", request.Method))
}

func (h *NestedRequestHandler) Describe() map[string]interface{} {
	tmp := map[string]interface{}{}
	for _, processor := range h.handlers {
		for key, val := range processor.Describe() {
			tmp[key] = val
		}
	}
	res := map[string]interface{}{h.Path(): tmp}
	return res
}

func (h *NestedRequestHandler) AddProcessor(processor RequestHandler) {
	h.handlers = append(h.handlers, processor)
	h.multiplexer.Handle(processor.Path(), processor)
}

func NewNestedRequestHandler(path string, processors []RequestHandler, multiplexer Multiplexer) *NestedRequestHandler {
	return &NestedRequestHandler{ReachableRequestHandler: ReachableRequestHandler{path: path}, handlers: processors, multiplexer: multiplexer}
}
