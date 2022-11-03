package server

import "net/http"

type ReachableRequestHandler struct {
	path string
	http.Handler
}

func (h *ReachableRequestHandler) Path() string {
	return h.path
}
