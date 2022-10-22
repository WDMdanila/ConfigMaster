package parameters

import (
	"config_master/internal/utils"
	"math/rand"
	"sync"
)

type SynchronizableSequenceParameter[T any] struct {
	SimpleParameter[T]
	mutex sync.Mutex
}

type RandomParameter struct {
	NamedParameter
	min int
	max int
}

func (parameter *RandomParameter) Set(i []byte) {
	//TODO implement me
	panic("implement me")
}

type ArithmeticSequenceParameter struct {
	SynchronizableSequenceParameter[float64]
	increment float64
}

type GeometricSequenceParameter struct {
	SynchronizableSequenceParameter[float64]
	multiplier float64
}

func (parameter *RandomParameter) GetAsJSON() []byte {
	return utils.GetAsJSON(parameter.name, rand.Intn(parameter.max-parameter.min)+parameter.min)
}

func NewRandomParameter(name string, min int, max int) Parameter {
	return &RandomParameter{NamedParameter: NamedParameter{name: name}, min: min, max: max}
}

func (parameter *ArithmeticSequenceParameter) GetAsJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.GetAsJSON()
}

func (parameter *ArithmeticSequenceParameter) update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter(name string, value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter[float64]{SimpleParameter: SimpleParameter[float64]{NamedParameter{name: name}, value}},
		increment:                       increment,
	}
}

func (parameter *GeometricSequenceParameter) GetAsJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.GetAsJSON()
}

func (parameter *GeometricSequenceParameter) update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter(name string, value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter[float64]{SimpleParameter: SimpleParameter[float64]{NamedParameter{name: name}, value}},
		multiplier:                      multiplier,
	}
}
