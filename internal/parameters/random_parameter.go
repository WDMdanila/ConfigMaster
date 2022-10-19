package parameters

import (
	"encoding/json"
	"math/rand"
)

type RandomParameter struct {
	min   int
	max   int
	Value int `json:"value"`
}

func (parameter *RandomParameter) AsJSON() []byte {
	parameter.Value = rand.Intn(parameter.max-parameter.min) + parameter.min
	jsonBytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func NewRandomParameter(min int, max int) Parameter {
	return &RandomParameter{min: min, max: max}
}
