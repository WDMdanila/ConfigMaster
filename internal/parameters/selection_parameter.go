package parameters

import (
	"math/rand"
	"sync"
)

type SelectionParameter struct {
	SimpleParameter[[]interface{}]
	mutex sync.Mutex
}

func (p *SelectionParameter) Extend(value interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.value = append(p.value, value)
}

type RandomSelectionParameter struct {
	SelectionParameter
}

type SequentialSelectionParameter struct {
	SelectionParameter
	index int
}

func (p *RandomSelectionParameter) Value() interface{} {
	return p.value[rand.Intn(len(p.value))]
}

func (p *RandomSelectionParameter) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: map[string]interface{}{
		"values":         p.value,
		"parameter_type": "random selection"},
	}
}

func (p *SequentialSelectionParameter) Set(data interface{}) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.index = 0
	return p.SimpleParameter.Set(data)
}

func (p *SequentialSelectionParameter) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: map[string]interface{}{
		"values":         p.value,
		"parameter_type": "sequential selection"},
	}
}

func (p *SequentialSelectionParameter) Value() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	defer p.updateIndex()
	return p.value[p.index]
}

func (p *SequentialSelectionParameter) updateIndex() {
	p.index = (p.index + 1) % len(p.value)
}

func NewRandomSelectionParameter(name string, options []interface{}) Parameter {
	return &RandomSelectionParameter{SelectionParameter: SelectionParameter{SimpleParameter: SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, value: options}}}
}

func NewSequentialSelectionParameter(name string, options []interface{}) Parameter {
	return &SequentialSelectionParameter{SelectionParameter: SelectionParameter{SimpleParameter: SimpleParameter[[]interface{}]{NamedParameter: NamedParameter{name: name}, value: options}}}
}
