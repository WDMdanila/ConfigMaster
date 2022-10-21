package parameters

import (
	"config_master/internal/utils"
	"math/rand"
	"sync"
)

type SelectionParameter[T any] struct {
	NamedParameter
	options []T
}

type RandomSelectionParameter[T any] struct {
	SelectionParameter[T]
}

func (parameter *RandomSelectionParameter[T]) Set(i interface{}) {
	//TODO implement me
	panic("implement me")
}

type SequentialSelectionParameter[T any] struct {
	SelectionParameter[T]
	index int
	mutex sync.Mutex
}

func (parameter *SequentialSelectionParameter[T]) Set(i interface{}) {
	//TODO implement me
	panic("implement me")
}

func (parameter *RandomSelectionParameter[T]) GetAsJSON() []byte {
	return utils.GetAsJSON(parameter.name, parameter.options[rand.Intn(len(parameter.options))])
}

func (parameter *SequentialSelectionParameter[T]) GetAsJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.updateIndex()
	return utils.GetAsJSON(parameter.name, parameter.options[parameter.index])
}

func (parameter *SequentialSelectionParameter[T]) updateIndex() {
	parameter.index = (parameter.index + 1) % len(parameter.options)
}

func NewRandomSelectionParameter[T any](name string, options []T) Parameter {
	return &RandomSelectionParameter[T]{SelectionParameter: SelectionParameter[T]{NamedParameter: NamedParameter{name: name}, options: options}}
}

func NewSequentialSelectionParameter[T any](name string, options []T) Parameter {
	return &SequentialSelectionParameter[T]{SelectionParameter: SelectionParameter[T]{NamedParameter: NamedParameter{name: name}, options: options}}
}
