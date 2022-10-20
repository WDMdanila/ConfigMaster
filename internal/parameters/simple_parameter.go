package parameters

import (
	"encoding/json"
)

type SimpleParameter[T any] struct {
	Value T `json:"value"`
}

func (parameter *SimpleParameter[T]) ToJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

type JSONParameter = SimpleParameter[interface{}]

func NewSimpleParameter[T any](data T) Parameter {
	return &SimpleParameter[T]{data}
}

func NewJSONParameter(data interface{}) Parameter {
	switch value := data.(type) {
	case []byte:
		var unpacked interface{}
		err := json.Unmarshal(value, &unpacked)
		if err != nil {
			panic(err)
		}
		return &JSONParameter{unpacked}
	default:
		return &JSONParameter{data}
	}
}
