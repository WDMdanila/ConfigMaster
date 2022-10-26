package parameters

import (
	"math/rand"
	"sync"
)

type SelectionParameter struct {
	SimpleParameter[[]interface{}]
}

type RandomSelectionParameter struct {
	SelectionParameter
}

type SequentialSelectionParameter struct {
	SelectionParameter
	index int
	mutex sync.Mutex
}

func (parameter *RandomSelectionParameter) Value() interface{} {
	return parameter.value[rand.Intn(len(parameter.value))]
}

func (parameter *RandomSelectionParameter) Describe() map[string]interface{} {
	return map[string]interface{}{parameter.name: map[string]interface{}{
		"values":         parameter.value,
		"parameter_type": "random selection"},
	}
}

func (parameter *SequentialSelectionParameter) Set(data interface{}) error {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	parameter.index = 0
	return parameter.SimpleParameter.Set(data)
}

func (parameter *SequentialSelectionParameter) Describe() map[string]interface{} {
	return map[string]interface{}{parameter.name: map[string]interface{}{
		"values":         parameter.value,
		"parameter_type": "sequential selection"},
	}
}

func (parameter *SequentialSelectionParameter) Value() interface{} {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.updateIndex()
	return parameter.value[parameter.index]
}

func (parameter *SequentialSelectionParameter) updateIndex() {
	parameter.index = (parameter.index + 1) % len(parameter.value)
}

func NewRandomSelectionParameter(name string, options []interface{}) Parameter {
	return &RandomSelectionParameter{SelectionParameter: SelectionParameter{SimpleParameter: SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, value: options}}}
}

func NewSequentialSelectionParameter(name string, options []interface{}) Parameter {
	return &SequentialSelectionParameter{SelectionParameter: SelectionParameter{SimpleParameter: SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, value: options}}}
}
