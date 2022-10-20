package parameters

import (
	"math/rand"
)

type RandomParameter struct {
	SimpleParameter[int]
	min int
	max int
}

func (parameter *RandomParameter) ToJSON() []byte {
	parameter.Value = rand.Intn(parameter.max-parameter.min) + parameter.min
	return parameter.SimpleParameter.ToJSON()
}

func NewRandomParameter(min int, max int) Parameter {
	return &RandomParameter{min: min, max: max}
}
