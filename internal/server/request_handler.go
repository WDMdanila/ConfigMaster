package server

import (
	"net/http"
)

type RequestHandler interface {
	http.Handler
	Path() string
	Process(*http.Request) map[string]interface{}
}
