package parameters

import (
	"math/rand"
	"sync"
)

type SequenceParameter interface {
	Parameter
	update()
}

type SynchronizableSequenceParameter[T any] struct {
	SimpleParameter[T]
	mutex sync.Mutex
}

type RandomParameter struct {
	SynchronizableSequenceParameter[int]
	min int
	max int
}

type ArithmeticSequenceParameter struct {
	SynchronizableSequenceParameter[float64]
	increment float64
}

type GeometricSequenceParameter struct {
	SynchronizableSequenceParameter[float64]
	multiplier float64
}

func (parameter *RandomParameter) ToJSON() []byte {
	parameter.mutex.Lock()
	parameter.update()
	defer parameter.mutex.Unlock()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *RandomParameter) update() {
	parameter.Value = rand.Intn(parameter.max-parameter.min) + parameter.min
}

func NewRandomParameter(min int, max int) Parameter {
	return &RandomParameter{min: min, max: max}
}

func (parameter *ArithmeticSequenceParameter) ToJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *ArithmeticSequenceParameter) update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter(value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter[float64]{SimpleParameter: SimpleParameter[float64]{value}},
		increment:                       increment,
	}
}

func (parameter *GeometricSequenceParameter) ToJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.ToJSON()
}

func (parameter *GeometricSequenceParameter) update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter(value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter[float64]{SimpleParameter: SimpleParameter[float64]{value}},
		multiplier:                      multiplier,
	}
}
