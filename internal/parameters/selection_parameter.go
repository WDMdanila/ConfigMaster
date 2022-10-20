package parameters

import (
	"math/rand"
	"sync"
)

type SelectionParameter[T any] struct {
	SimpleParameter[T]
	options []T
}

type RandomSelectionParameter[T any] struct {
	SelectionParameter[T]
}

type SequentialSelectionParameter[T any] struct {
	SelectionParameter[T]
	index int
	mutex sync.Mutex
}

func (parameter *RandomSelectionParameter[T]) ToJSON() []byte {
	parameter.Value = parameter.options[rand.Intn(len(parameter.options))]
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *SequentialSelectionParameter[T]) ToJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	parameter.Value = parameter.options[parameter.index]
	parameter.updateIndex()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *SequentialSelectionParameter[T]) updateIndex() {
	parameter.index = (parameter.index + 1) % len(parameter.options)
}

func NewRandomSelectionParameter[T any](options []T) Parameter {
	return &RandomSelectionParameter[T]{SelectionParameter: SelectionParameter[T]{options: options}}
}

func NewSequentialSelectionParameter[T any](options []T) Parameter {
	return &SequentialSelectionParameter[T]{SelectionParameter: SelectionParameter[T]{options: options}}
}
