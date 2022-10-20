package parameters

import (
	"encoding/json"
	"math/rand"
)

type SelectionParameter[T any] struct {
	options []T
	Value   T `json:"value"`
}

type RandomSelectionParameter[T any] struct {
	SelectionParameter[T]
}

type SequentialSelectionParameter[T any] struct {
	SelectionParameter[T]
	index int
}

func (parameter *RandomSelectionParameter[T]) ToJSON() []byte {
	parameter.Value = parameter.options[rand.Intn(len(parameter.options))]
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (parameter *SequentialSelectionParameter[T]) ToJSON() []byte {
	parameter.Value = parameter.options[parameter.index]
	parameter.updateIndex()
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	return jsonBytes
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
