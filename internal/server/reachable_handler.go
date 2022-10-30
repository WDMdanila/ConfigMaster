package server

type ReachableRequestHandler struct {
	path string
}

func (h *ReachableRequestHandler) Path() string {
	return h.path
}
