package parameters

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ParameterReader interface {
	Read() map[string]Parameter
}

type JSONParameterReader struct {
	filePath string
}

func (parameterReader *JSONParameterReader) Read() map[string]Parameter {
	data := parseJSONFile(parameterReader.filePath)
	res := map[string]Parameter{}
	for key, element := range data {
		res[key] = parseParameter(element)
	}
	return res
}

func NewJSONParameterReader(filePath string) ParameterReader {
	return &JSONParameterReader{filePath: filePath}
}

func parseJSONFile(filePath string) map[string]map[string]interface{} {
	rawData := readAllFromFile(filePath)
	var data map[string]map[string]interface{}
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic(err)
	}
	return data
}

func readAllFromFile(filePath string) []byte {
	jsonFile := openFile(filePath)
	defer closeFile(jsonFile)
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Panic(err)
	}
	return byteValue
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("error during closing file: %v", err)
	}
}

func parseParameter(element map[string]interface{}) Parameter {
	elem := parseString(element, "type")
	f := getParserFunc(elem)
	return f(element)
}

func parseString(element map[string]interface{}, name string) string {
	switch elem := element[name].(type) {
	case string:
		return elem
	default:
		panic(fmt.Errorf(`could not parse "%v", got type "%T"`, elem, elem))
	}
}

func getParserFunc(elem string) func(map[string]interface{}) Parameter {
	f, ok := parameterTypeParserMap[elem]
	if !ok {
		panic(fmt.Errorf("\"%v\" parameter type is unknown\n", elem))
	}
	return f
}

var parameterTypeParserMap = map[string]func(map[string]interface{}) Parameter{
	"number":               ParseSimpleParameter[float64],
	"bool":                 ParseSimpleParameter[bool],
	"string":               ParseSimpleParameter[string],
	"json":                 ParseJSONParameter,
	"random":               ParseRandomParameter,
	"arithmetic sequence":  ParseArithmeticSequenceParameter,
	"geometric sequence":   ParseGeometricSequenceParameter,
	"sequential selection": ParseSequentialSelectionParameter,
	"random selection":     ParseRandomSelectionParameter,
}

func retrieveFloat(element map[string]interface{}, name string) float64 {
	switch elem := element[name].(type) {
	case float64:
		return elem
	default:
		panic(fmt.Errorf(`could not parse "%v", got type "%T"`, elem, elem))
	}
}

func ParseSimpleParameter[T any](element map[string]interface{}) Parameter {
	switch elem := element["value"].(type) {
	case T:
		return NewSimpleParameter(elem)
	default:
		panic(fmt.Errorf(`could not parse "%v", got type "%T"`, elem, elem))
	}
}

func ParseJSONParameter(element map[string]interface{}) Parameter {
	switch elem := element["value"].(type) {
	case map[string]interface{}:
		return NewJSONParameter(elem)
	default:
		panic(fmt.Errorf(`could not parse "%v", got type "%T"`, elem, elem))
	}
}

func ParseRandomParameter(element map[string]interface{}) Parameter {
	min := int(retrieveFloat(element, "min"))
	max := int(retrieveFloat(element, "max"))
	return NewRandomParameter(min, max)
}

func ParseArithmeticSequenceParameter(element map[string]interface{}) Parameter {
	value := retrieveFloat(element, "value")
	increment := retrieveFloat(element, "increment")
	return NewArithmeticSequenceParameter(value, increment)
}

func ParseGeometricSequenceParameter(element map[string]interface{}) Parameter {
	value := retrieveFloat(element, "value")
	multiplier := retrieveFloat(element, "multiplier")
	return NewGeometricSequenceParameter(value, multiplier)
}

func ParseSequentialSelectionParameter(element map[string]interface{}) Parameter {
	selection := parseArray(element, "values")
	return NewSequentialSelectionParameter(selection)
}

func parseArray(element map[string]interface{}, name string) []interface{} {
	switch values := element[name].(type) {
	case []interface{}:
		return values
	default:
		panic(fmt.Errorf(`expected array as "%v", got type: %T`, name, values))
	}
}

func ParseRandomSelectionParameter(element map[string]interface{}) Parameter {
	selection := parseArray(element, "values")
	return NewRandomSelectionParameter(selection)
}
