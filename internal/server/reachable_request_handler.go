package server

type ReachableRequestHandler struct {
	path string
}

func (handler *ReachableRequestHandler) Path() string {
	return handler.path
}
