package parameters

import (
	"config_master/internal/utils"
	"encoding/json"
)

type NamedParameter struct {
	name string
}

type SimpleParameter struct {
	NamedParameter
	Value interface{}
}

func (parameter *SimpleParameter) GetAsJSON() []byte {
	tmp := map[string]interface{}{parameter.name: parameter.Value}
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (parameter *NamedParameter) Name() string {
	return parameter.name
}

func (parameter *SimpleParameter) Set(data []byte) {
	parameter.Value = utils.ExtractFromJSON[interface{}](data, "value")
}

func NewSimpleParameter(name string, data interface{}) Parameter {
	switch value := data.(type) {
	case []byte:
		return &SimpleParameter{NamedParameter: NamedParameter{name: name}, Value: utils.DecodeJSON[interface{}](value)}
	default:
		return &SimpleParameter{NamedParameter: NamedParameter{name: name}, Value: value}
	}
}
