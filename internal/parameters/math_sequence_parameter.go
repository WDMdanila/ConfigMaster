package parameters

import (
	"encoding/json"
)

type ArithmeticSequenceParameter struct {
	SimpleParameter[float64]
	increment float64
}

type GeometricSequenceParameter struct {
	SimpleParameter[float64]
	multiplier float64
}

func (parameter *ArithmeticSequenceParameter) ToJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	parameter.Update()
	return jsonBytes
}

func (parameter *ArithmeticSequenceParameter) Update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter(value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{SimpleParameter: SimpleParameter[float64]{value}, increment: increment}
}

func (parameter *GeometricSequenceParameter) ToJSON() []byte {
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	parameter.Update()
	return jsonBytes
}

func (parameter *GeometricSequenceParameter) Update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter(value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{SimpleParameter: SimpleParameter[float64]{value}, multiplier: multiplier}
}
