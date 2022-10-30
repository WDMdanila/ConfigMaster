package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type NestedRequestHandler struct {
	ReachableRequestHandler
	processors []RequestHandler
}

func (h *NestedRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %v request to: %v\n", request.Method, request.RequestURI)
	writer.Header().Set("Content-Type", "application/json")
	data := h.GetResponse(request)
	_, err := writer.Write(data)
	if err != nil {
		log.Printf("could not respond to: %v, error: %v", request.RemoteAddr, err)
		return
	}
	log.Printf("responded to: %v with: %v\n", request.RemoteAddr, string(data))
}

func (h *NestedRequestHandler) GetResponse(request *http.Request) []byte {
	result := h.Process(request)
	data, _ := json.Marshal(result)
	return data
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

func NewNestedRequestHandler(path string, processors []RequestHandler) *NestedRequestHandler {
	return &NestedRequestHandler{ReachableRequestHandler: ReachableRequestHandler{path: path}, processors: processors}
}
