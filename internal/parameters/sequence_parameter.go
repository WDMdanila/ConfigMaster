package parameters

import (
	"config_master/internal/utils"
	"encoding/json"
	"math/rand"
	"sync"
)

type SynchronizableSequenceParameter struct {
	NamedParameter
	Value float64
	mutex sync.Mutex
}

type RandomParameter struct {
	NamedParameter
	min int
	max int
}

func (parameter *RandomParameter) Set(i []byte) {
	parameter.min = int(utils.ExtractFromJSON[float64](i, "min"))
	parameter.max = int(utils.ExtractFromJSON[float64](i, "max"))
}

type ArithmeticSequenceParameter struct {
	SynchronizableSequenceParameter
	increment float64
}

func (parameter *ArithmeticSequenceParameter) Set(bytes []byte) {
	//TODO implement me
	panic("implement me")
}

type GeometricSequenceParameter struct {
	SynchronizableSequenceParameter
	multiplier float64
}

func (parameter *GeometricSequenceParameter) Set(bytes []byte) {
	//TODO implement me
	panic("implement me")
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
	tmp := map[string]interface{}{parameter.name: parameter.Value}
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (parameter *ArithmeticSequenceParameter) update() {
	parameter.Value += parameter.increment
}

func NewArithmeticSequenceParameter(name string, value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{NamedParameter: NamedParameter{name: name}, Value: value},
		increment:                       increment,
	}
}

func (parameter *GeometricSequenceParameter) GetAsJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	tmp := map[string]interface{}{parameter.name: parameter.Value}
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (parameter *GeometricSequenceParameter) update() {
	parameter.Value *= parameter.multiplier
}

func NewGeometricSequenceParameter(name string, value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{NamedParameter: NamedParameter{name: name}, Value: value},
		multiplier:                      multiplier,
	}
}
