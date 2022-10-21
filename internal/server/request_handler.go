package server

import (
	"net/http"
)

type RequestHandler interface {
	http.Handler
	Path() string
}

type Processor interface {
	Process(*http.Request) []byte
}
