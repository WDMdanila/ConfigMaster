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

func (p *RandomParameter) Set(data interface{}) error {
	switch value := data.(type) {
	case map[string]float64:
		if val, ok := value["min"]; ok {
			p.min = int(val)
		}
		if val, ok := value["max"]; ok {
			p.max = int(val)
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", p.name, value, value, map[string]float64{})
	}
}

type ArithmeticSequenceParameter struct {
	SynchronizableSequenceParameter
	increment float64
}

func (p *ArithmeticSequenceParameter) Set(data interface{}) error {
	switch fieldData := data.(type) {
	case map[string]float64:
		if val, ok := fieldData["increment"]; ok {
			p.increment = val
		}
		if val, ok := fieldData["value"]; ok {
			p.value = val
		}
		return nil
	case map[string]interface{}:
		val, err := utils.ParseFloat(fieldData, "increment")
		if err == nil {
			p.increment = val
		}
		val, err = utils.ParseFloat(fieldData, "value")
		if err == nil {
			p.value = val
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", p.name, fieldData, fieldData, float64(0))
	}
}

type GeometricSequenceParameter struct {
	SynchronizableSequenceParameter
	multiplier float64
}

func (p *GeometricSequenceParameter) Set(data interface{}) error {
	switch fieldData := data.(type) {
	case map[string]float64:
		if val, ok := fieldData["multiplier"]; ok {
			p.multiplier = val
		}
		if val, ok := fieldData["value"]; ok {
			p.value = val
		}
		return nil
	case map[string]interface{}:
		val, err := utils.ParseFloat(fieldData, "multiplier")
		if err == nil {
			p.multiplier = val
		}
		val, err = utils.ParseFloat(fieldData, "value")
		if err == nil {
			p.value = val
		}
		return nil
	default:
		return fmt.Errorf("failed to set %v to %v due to type mismatch (got %T, expected %T)", p.name, fieldData, fieldData, float64(0))
	}
}

func (p *RandomParameter) Value() interface{} {
	return rand.Intn(p.max-p.min) + p.min
}

func (p *RandomParameter) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: map[string]interface{}{
		"min":            p.min,
		"max":            p.max,
		"parameter_type": "random"},
	}
}

func NewRandomParameter(name string, min int, max int) Parameter {
	return &RandomParameter{NamedParameter: NamedParameter{name: name}, min: min, max: max}
}

func (p *ArithmeticSequenceParameter) Value() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	defer p.update()
	return p.SimpleParameter.Value()
}

func (p *ArithmeticSequenceParameter) update() {
	p.value += p.increment
}

func (p *ArithmeticSequenceParameter) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: map[string]interface{}{
		"value":          p.value,
		"increment":      p.increment,
		"parameter_type": "arithmetic sequence"},
	}
}

func NewArithmeticSequenceParameter(name string, value float64, increment float64) Parameter {
	return &ArithmeticSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{SimpleParameter: SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, value: value}},
		increment:                       increment,
	}
}

func (p *GeometricSequenceParameter) Value() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	defer p.update()
	return p.SimpleParameter.Value()
}

func (p *GeometricSequenceParameter) update() {
	p.value *= p.multiplier
}

func (p *GeometricSequenceParameter) Describe() map[string]interface{} {
	return map[string]interface{}{p.name: map[string]interface{}{
		"value":          p.value,
		"multiplier":     p.multiplier,
		"parameter_type": "geometric sequence"},
	}
}

func NewGeometricSequenceParameter(name string, value float64, multiplier float64) Parameter {
	return &GeometricSequenceParameter{
		SynchronizableSequenceParameter: SynchronizableSequenceParameter{SimpleParameter: SimpleParameter[float64]{NamedParameter: NamedParameter{name: name}, value: value}},
		multiplier:                      multiplier,
	}
}
