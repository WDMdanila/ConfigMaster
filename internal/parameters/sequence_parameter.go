package parameters

import (
	"config_master/internal/utils"
	"fmt"
	"math/rand"
	"sync"
)

type SynchronizableSequenceParameter struct {
	SimpleParameter[float64]
	mutex sync.Mutex
}

type RandomParameter struct {
	NamedParameter
	min int
	max int
}

func (parameter *RandomParameter) Set(data interface{}) error {
	switch value := data.(type) {
	case map[string]float64:
		if val, ok := value["min"]; ok {
			parameter.min = int(val)
		}
		if val, ok := value["max"]; ok {
			parameter.max = int(val)
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", parameter.name, value, value, map[string]float64{})
	}
}

type ArithmeticSequenceParameter struct {
	SynchronizableSequenceParameter
	increment float64
}

func (parameter *ArithmeticSequenceParameter) Set(data interface{}) error {
	switch fieldData := data.(type) {
	case map[string]float64:
		if val, ok := fieldData["increment"]; ok {
			parameter.increment = val
		}
		if val, ok := fieldData["value"]; ok {
			parameter.value = val
		}
		return nil
	case map[string]interface{}:
		val, err := utils.ParseFloat(fieldData, "increment")
		if err == nil {
			parameter.increment = val
		}
		val, err = utils.ParseFloat(fieldData, "value")
		if err == nil {
			parameter.value = val
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", parameter.name, fieldData, fieldData, float64(0))
	}
}

type GeometricSequenceParameter struct {
	SynchronizableSequenceParameter
	multiplier float64
}

func (parameter *GeometricSequenceParameter) Set(data interface{}) error {
	switch fieldData := data.(type) {
	case map[string]float64:
		if val, ok := fieldData["multiplier"]; ok {
			parameter.multiplier = val
		}
		if val, ok := fieldData["value"]; ok {
			parameter.value = val
		}
		return nil
	case map[string]interface{}:
		val, err := utils.ParseFloat(fieldData, "multiplier")
		if err == nil {
			parameter.multiplier = val
		}
		val, err = utils.ParseFloat(fieldData, "value")
		if err == nil {
			parameter.value = val
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", parameter.name, fieldData, fieldData, float64(0))
	}
}

func (parameter *RandomParameter) Value() interface{} {
	return rand.Intn(parameter.max-parameter.min) + parameter.min
}

func NewRandomParameter(name string, min int, max int) Parameter {
	return &RandomParameter{NamedParameter: NamedParameter{name: name}, min: min, max: max}
}

func (parameter *ArithmeticSequenceParameter) Value() interface{} {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.Value()
}

func (parameter *ArithmeticSequenceParameter) update() {
	parameter.value += parameter.increment
}

func NewArithmeticSequenceParameter(name string, value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{SimpleParameter: SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, value: value}},
		increment:                       increment,
	}
}

func (parameter *GeometricSequenceParameter) Value() interface{} {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.update()
	return parameter.SimpleParameter.Value()
}

func (parameter *GeometricSequenceParameter) update() {
	parameter.value *= parameter.multiplier
}

func NewGeometricSequenceParameter(name string, value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{SimpleParameter: SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, value: value}},
		multiplier:                      multiplier,
	}
}
