package server

import (
	"config_master/internal/parameters"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type TimestampHandler struct {
	delay uint64
	mutex sync.Mutex
	parameters.SimpleParameter[uint64]
}

func (handler *TimestampHandler) Process(request *http.Request) ([]byte, error) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()
	if handler.Value != 0 {
		return handler.ToJSON(), nil
	}
	parsedUint, err := parseRequest(request)
	if err != nil {
		return nil, err
	}
	handler.Value = parsedUint + handler.delay
	return handler.ToJSON(), nil
}

func parseRequest(request *http.Request) (uint64, error) {
	timestampString, err := getTimestampFromUrl(request.URL)
	if err != nil {
		log.Printf("could not process timestamp, error: %v\n", err)
		return 0, err
	}
	parsedUint, err := strconv.ParseUint(timestampString, 10, 64)
	if err != nil {
		log.Printf("could not process timestamp, error: %v\n", err)
		return 0, err
	}
	return parsedUint, nil
}

func getTimestampFromUrl(url *url.URL) (string, error) {
	rawValue := url.Query().Get("timestamp")
	if rawValue == "" {
		return "", errors.New(`"timestamp" parameter missing from url query`)
	}
	return rawValue, nil
}

func NewTimestampHandler(path string, delayNs uint64) RequestHandler {
	return &DefaultRequestHandler{path: path, Processor: &TimestampHandler{delay: delayNs}}
}
