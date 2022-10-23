package parameters

import "fmt"

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
		panic(fmt.Errorf(`could not parse "%v", got type "%T"`, elem, elem))
	}
}

func retrieveInt(element map[string]interface{}, name string) int {
	var value int
	switch elem := element[name].(type) {
	case float64:
		value = int(elem)
	default:
		err := fmt.Sprintf(`could not parse "%v", got type "%T"`, elem, elem)
		panic(err)
	}
	return value
}
