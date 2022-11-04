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

func (p *NamedParameter) Name() string {
	return p.name
}

func (p *SimpleParameter[T]) Value() interface{} {
	return p.value
}

func (p *SimpleParameter[T]) Set(data interface{}) error {
	switch value := data.(type) {
	case T:
		p.value = value
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", p.name, value, value, p.value)
	}
}

func (p *SimpleParameter[T]) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: p.value}
}

func NewSimpleParameter(name string, data interface{}) *SimpleParameter[interface{}] {
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
