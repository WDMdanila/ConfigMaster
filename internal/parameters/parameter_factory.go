package parameters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var parameterTypeMap = map[string]func(string, map[string]interface{}) Parameter{
	"sequential selection": newSequentialSelectionParameter,
	"random selection":     newRandomSelectionParameter,
	"arithmetic sequence":  newArithmeticSequenceParameter,
	"geometric sequence":   newGeometricSequenceParameter,
	"random":               newRandomParameter,
}

func FromJSON(name string, json map[string]interface{}, strictType bool) Parameter {
	if val, ok := json["parameter_type"]; ok {
		switch paramType := val.(type) {
		case string:
			return parameterTypeMap[paramType](name, json)
		}
	}
	if strictType {
		return NewSimpleStrictParameter(name, json)
	}
	return NewSimpleParameter(name, json)
}

func newSequentialSelectionParameter(name string, data map[string]interface{}) Parameter {
	if val, ok := data["source"]; ok {
		switch source := val.(type) {
		case string:
			if source == "files" {
				if val, ok := data["values"]; ok {
					switch values := val.(type) {
					case []interface{}:
						var vals []interface{}
						for _, file := range values {
							data, err := os.ReadFile(file.(string))
							if err != nil {
								panic(err)
							}
							data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
							var res interface{}
							err = json.Unmarshal(data, &res)
							if err != nil {
								panic(err)
							}
							vals = append(vals, res)
						}
						return NewSequentialSelectionParameter(name, vals)
					default:
						panic("shiet")
					}
				}
			}
		}
	}
	if val, ok := data["values"]; ok {
		switch values := val.(type) {
		case []interface{}:
			return NewSequentialSelectionParameter(name, values)
		default:
			panic(fmt.Errorf("%v parameter values are not a list", name))
		}
	}
	return NewSimpleParameter(name, data)
}

func newRandomSelectionParameter(name string, data map[string]interface{}) Parameter {
	if val, ok := data["values"]; ok {
		switch values := val.(type) {
		case []interface{}:
			return NewRandomSelectionParameter(name, values)
		default:
			panic(fmt.Errorf("%v parameter values are not a list", name))
		}
	}
	return NewSimpleParameter(name, data)
}

func newArithmeticSequenceParameter(name string, data map[string]interface{}) Parameter {
	value := retrieveFloat(data, "value")
	increment := retrieveFloat(data, "increment")
	return NewArithmeticSequenceParameter(name, value, increment)
}

func newGeometricSequenceParameter(name string, data map[string]interface{}) Parameter {
	value := retrieveFloat(data, "value")
	multiplier := retrieveFloat(data, "multiplier")
	return NewGeometricSequenceParameter(name, value, multiplier)
}

func newRandomParameter(name string, data map[string]interface{}) Parameter {
	min := retrieveInt(data, "min")
	max := retrieveInt(data, "max")
	return NewRandomParameter(name, min, max)
}

func retrieveFloat(element map[string]interface{}, name string) float64 {
	switch elem := element[name].(type) {
	case float64:
		return elem
	default:
		panic(fmt.Errorf("could not convert %T to float", elem))
	}
}

func retrieveInt(element map[string]interface{}, name string) int {
	switch elem := element[name].(type) {
	case float64:
		return int(elem)
	default:
		panic(fmt.Errorf("could not convert %T to int", elem))
	}
}
