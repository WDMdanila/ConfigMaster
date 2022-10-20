package parameters

import "sync"

type ArithmeticSequenceParameter struct {
	SimpleParameter[float64]
	increment float64
	mutex     sync.Mutex
}

type GeometricSequenceParameter struct {
	SimpleParameter[float64]
	multiplier float64
	mutex      sync.Mutex
}

func (parameter *ArithmeticSequenceParameter) ToJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.Update()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *ArithmeticSequenceParameter) Update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter(value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{SimpleParameter: SimpleParameter[float64]{value}, increment: increment}
}

func (parameter *GeometricSequenceParameter) ToJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.Update()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *GeometricSequenceParameter) Update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter(value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{SimpleParameter: SimpleParameter[float64]{value}, multiplier: multiplier}
}
