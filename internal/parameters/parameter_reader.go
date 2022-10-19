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
		log.Fatal(err)
	}
	return data
}

func readAllFromFile(filePath string) []byte {
	jsonFile := openFile(filePath)
	defer closeFile(jsonFile)
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	return byteValue
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
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
		err := fmt.Sprintf(`could not parse "%v", got type "%T"`, elem, elem)
		panic(err)
	}
}

func getParserFunc(elem string) func(map[string]interface{}) Parameter {
	f, ok := parameterTypeParserMap[elem]
	if !ok {
		log.Fatalf(`"%v" parameter type is unknown\n`, elem)
	}
	return f
}

var parameterTypeParserMap = map[string]func(map[string]interface{}) Parameter{
	"int":                  ParseIntParameter,
	"float":                ParseSimpleParameter[float64],
	"bool":                 ParseSimpleParameter[bool],
	"string":               ParseSimpleParameter[string],
	"json":                 ParseJSONParameter,
	"random":               ParseRandomParameter,
	"arithmetic sequence":  ParseArithmeticSequenceParameter,
	"geometric sequence":   ParseGeometricSequenceParameter,
	"sequential selection": ParseSequentialSelectionParameter,
	"random selection":     ParseRandomSelectionParameter,
}

func ParseIntParameter(element map[string]interface{}) Parameter {
	value := retrieveInt(element, "value")
	return &SimpleParameter[int]{value}
}

func retrieveInt(element map[string]interface{}, name string) int {
	var value int
	switch elem := element[name].(type) {
	case int:
		value = elem
	case float64:
		verifyIntegral(elem)
		value = int(elem)
	default:
		err := fmt.Sprintf(`could not parse "%v", got type "%T"`, elem, elem)
		panic(err)
	}
	return value
}

func verifyIntegral(elem float64) {
	if elem != float64(int(elem)) {
		err := fmt.Sprintf(`tried to parse float value "%v" as int`, elem)
		panic(err)
	}
}

func ParseSimpleParameter[T any](element map[string]interface{}) Parameter {
	switch elem := element["value"].(type) {
	case T:
		return &SimpleParameter[T]{elem}
	default:
		err := fmt.Sprintf(`could not parse "%v", got type "%T"`, elem, elem)
		panic(err)
	}
}

func ParseJSONParameter(element map[string]interface{}) Parameter {
	switch elem := element["value"].(type) {
	case map[string]interface{}:
		data := toBytes(elem)
		return NewJSONParameter(data)
	default:
		err := fmt.Sprintf(`could not parse "%v", got type "%T"`, elem, elem)
		panic(err)
	}
}

func toBytes(elem map[string]interface{}) []byte {
	data, err := json.Marshal(elem)
	if err != nil {
		log.Fatalf(`could not parse "%v", error: %v\n`, elem, err)
	}
	return data
}

func ParseRandomParameter(element map[string]interface{}) Parameter {
	min := retrieveInt(element, "min")
	max := retrieveInt(element, "max")
	return &RandomParameter{min: min, max: max}
}

func ParseArithmeticSequenceParameter(element map[string]interface{}) Parameter {
	value := retrieveInt(element, "value")
	increment := retrieveInt(element, "increment")
	return &ArithmeticSequenceParameter[int]{value, increment}
}

func ParseGeometricSequenceParameter(element map[string]interface{}) Parameter {
	value := retrieveInt(element, "value")
	multiplier := retrieveInt(element, "multiplier")
	return &GeometricSequenceParameter[int]{value, multiplier}
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
		errString := fmt.Sprintf(`expected array as "%v", got type: %T`, name, values)
		panic(errString)
	}
}

func ParseRandomSelectionParameter(element map[string]interface{}) Parameter {
	selection := parseArray(element, "values")
	return NewRandomSelectionParameter(selection)
}
