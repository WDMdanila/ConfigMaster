package parameters

import (
	"fmt"
)

type NamedParameter struct {
	name string
}

type SimpleParameter[T any] struct {
	NamedParameter
	value T
}

func (parameter *NamedParameter) Name() string {
	return parameter.name
}

func (parameter *SimpleParameter[T]) Value() interface{} {
	return parameter.value
}

func (parameter *SimpleParameter[T]) Set(data interface{}) error {
	switch value := data.(type) {
	case T:
		parameter.value = value
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", parameter.name, value, value, parameter.value)
	}
}

func (parameter *SimpleParameter[T]) Describe() map[string]interface{} {
	return map[string]interface{}{parameter.name: parameter.value}
}

func NewSimpleParameter(name string, data interface{}) Parameter {
	return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, value: data}
}

func NewSimpleStrictParameter(name string, data interface{}) Parameter {
	switch value := data.(type) {
	case float64:
		return &SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, value: value}
	case bool:
		return &SimpleParameter[bool]{NamedParameter: NamedParameter{name: name}, value: value}
	case string:
		return &SimpleParameter[string]{NamedParameter: NamedParameter{name: name}, value: value}
	case []interface{}:
		return &SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, value: value}
	case map[string]interface{}:
		return &SimpleParameter[map[string]interface{}]{NamedParameter: NamedParameter{name: name}, value: value}
	case []byte:
		return &SimpleParameter[[]byte]{NamedParameter: NamedParameter{name: name}, value: value}
	default:
		return &SimpleParameter[interface{}]{NamedParameter: NamedParameter{name: name}, value: value}
	}
}
