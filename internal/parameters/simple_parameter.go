package parameters

import (
	"config_master/internal/utils"
	"encoding/json"
	"fmt"
	"log"
)

type NamedParameter struct {
	name string
}

type SimpleParameter[T any] struct {
	NamedParameter
	Value T
}

func (parameter *SimpleParameter[T]) GetAsJSON() ([]byte, error) {
	tmp := map[string]interface{}{parameter.name: parameter.Value}
	return json.Marshal(tmp)
}

func (parameter *NamedParameter) Name() string {
	return parameter.name
}

func (parameter *SimpleParameter[T]) Set(data []byte) error {
	res, err := utils.ExtractFromJSON[T](data, "value")
	if err != nil {
		return fmt.Errorf("failed to set %v of type %T, error: %v", parameter.name, parameter.Value, err)
	}
	parameter.Value = res
	return nil
}

func NewSimpleParameter(name string, data interface{}) Parameter {
	switch value := data.(type) {
	case []byte:
		res, err := utils.DecodeJSON[interface{}](value)
		if err != nil {
			log.Panic(err)
		}
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: res}
	default:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	}
}

func NewSimpleStrictParameter(name string, data interface{}) Parameter {
	switch value := data.(type) {
	case float64:
		return &SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, Value: value}
	case bool:
		return &SimpleParameter[bool]{NamedParameter: NamedParameter{name: name}, Value: value}
	case string:
		return &SimpleParameter[string]{NamedParameter: NamedParameter{name: name}, Value: value}
	case []interface{}:
		return &SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	case map[string]interface{}:
		return &SimpleParameter[map[string]interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	case []byte:
		return parseRawJSON(name, value)
	default:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	}
}

func parseRawJSON(name string, value []byte) Parameter {
	val, err := utils.DecodeJSON[interface{}](value)
	if err != nil {
		log.Panic(err)
	}
	switch v := val.(type) {
	case map[string]interface{}:
		return &SimpleParameter[map[string]interface{}]{NamedParameter: NamedParameter{name: name}, Value: v}
	case []interface{}:
		return &SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, Value: v}
	default:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, Value: value}
	}
}
