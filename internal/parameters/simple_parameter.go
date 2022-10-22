package parameters

import (
	"config_master/internal/utils"
	"encoding/json"
)

type NamedParameter struct {
	name string
}

type SimpleParameter[T any] struct {
	NamedParameter
	Value T
}

func (parameter *SimpleParameter[T]) GetAsJSON() []byte {
	tmp := map[string]T{parameter.name: parameter.Value}
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (parameter *NamedParameter) Name() string {
	return parameter.name
}

func (parameter *SimpleParameter[T]) Set(data []byte) {
	parameter.Value = utils.ExtractFromJSON[T](data, "value")
}

func NewSimpleParameter(name string, data interface{}) Parameter {
	switch value := data.(type) {
	case float64:
		return &SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, Value: value}
	case bool:
		return &SimpleParameter[bool]{NamedParameter: NamedParameter{name: name}, Value: value}
	case string:
		return &SimpleParameter[string]{NamedParameter: NamedParameter{name: name}, Value: value}
	case []byte:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: utils.DecodeJSON[interface{}](value)}
	default:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	}
}
