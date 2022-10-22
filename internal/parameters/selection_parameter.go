package parameters

import (
	"config_master/internal/utils"
	"math/rand"
	"sync"
)

type SelectionParameter struct {
	NamedParameter
	Value []interface{}
}

func (parameter *SelectionParameter) Set(data []byte) error {
	res, err := utils.ExtractFromJSON[[]interface{}](data, "values")
	if err != nil {
		return err
	}
	parameter.Value = res
	return nil
}

type RandomSelectionParameter struct {
	SelectionParameter
}

type SequentialSelectionParameter struct {
	SelectionParameter
	index int
	mutex sync.Mutex
}

func (parameter *RandomSelectionParameter) GetAsJSON() []byte {
	return utils.GetAsJSON(parameter.name, parameter.Value[rand.Intn(len(parameter.Value))])
}

func (parameter *SequentialSelectionParameter) GetAsJSON() []byte {
	parameter.mutex.Lock()
	defer parameter.mutex.Unlock()
	defer parameter.updateIndex()
	return utils.GetAsJSON(parameter.name, parameter.Value[parameter.index])
}

func (parameter *SequentialSelectionParameter) updateIndex() {
	parameter.index = (parameter.index + 1) % len(parameter.Value)
}

func NewRandomSelectionParameter(name string, options []interface{}) Parameter {
	return &RandomSelectionParameter{SelectionParameter: SelectionParameter{NamedParameter: NamedParameter{name: name}, Value: options}}
}

func NewSequentialSelectionParameter(name string, options []interface{}) Parameter {
	return &SequentialSelectionParameter{SelectionParameter: SelectionParameter{NamedParameter: NamedParameter{name: name}, Value: options}}
}
