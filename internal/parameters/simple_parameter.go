package parameters

import (
	"encoding/json"
)

type SimpleParameter[T any] struct {
	Value T `json:"value"`
}

func (parameter *SimpleParameter[T]) AsJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

type JSONParameter = SimpleParameter[interface{}]

func NewJSONParameter(data []byte) Parameter {
	var unpacked interface{}
	err := json.Unmarshal(data, &unpacked)
	if err != nil {
		panic(err)
	}
	return &JSONParameter{unpacked}
}
