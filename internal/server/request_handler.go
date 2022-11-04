package server

import (
	"config_master/internal/utils"
	"io"
	"net/http"
)

type RequestHandler interface {
	http.Handler
	Path() string
	Describe() map[string]interface{}
}

func parseResponse(name string, value interface{}) []byte {
	res, _ := utils.GetAsJSON(name, value)
	return res
}

func extractData(request *http.Request) (map[string]interface{}, error) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	value, err := utils.DecodeJSON[map[string]interface{}](data)
	if err != nil {
		return nil, err
	}
	return value, nil
}
