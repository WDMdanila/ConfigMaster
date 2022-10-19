package parameters

import (
	"encoding/json"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type ArithmeticSequenceParameter[T Numbers] struct {
	Value     T `json:"value"`
	increment T
}

func (parameter *ArithmeticSequenceParameter[T]) AsJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	parameter.Update()
	return jsonBytes
}

func (parameter *ArithmeticSequenceParameter[T]) Update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter[T Numbers](value T, increment T) Parameter {
	return &ArithmeticSequenceParameter[T]{value, increment}
}

type GeometricSequenceParameter[T Numbers] struct {
	Value      T `json:"value"`
	multiplier T
}

func (parameter *GeometricSequenceParameter[T]) AsJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	parameter.Update()
	return jsonBytes
}

func (parameter *GeometricSequenceParameter[T]) Update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter[T Numbers](value T, multiplier T) Parameter {
	return &GeometricSequenceParameter[T]{value, multiplier}
}
