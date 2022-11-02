package server

import (
	"net/http"
)

type RequestHandler interface {
	http.Handler
	Path() string
	Describe() map[string]interface{}
}
