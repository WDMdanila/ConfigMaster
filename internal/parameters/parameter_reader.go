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
	filePath    string
	strictTypes bool
}

func (parameterReader *JSONParameterReader) Read() map[string]Parameter {
	data := parseJSONFile(parameterReader.filePath)
	res := map[string]Parameter{}
	for key, element := range data {
		parameter := parseParameter(key, element, parameterReader.strictTypes)
		res[parameter.Name()] = parameter
	}
	return res
}

func NewJSONParameterReader(filePath string, strictTypes bool) ParameterReader {
	return &JSONParameterReader{filePath: filePath, strictTypes: strictTypes}
}

func parseJSONFile(filePath string) map[string]interface{} {
	rawData := readAllFromFile(filePath)
	var data map[string]interface{}
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

func parseParameter(name string, element interface{}, strictType bool) Parameter {
	switch elem := element.(type) {
	case map[string]interface{}:
		return FromJSON(name, elem, strictType)
	default:
		if strictType {
			return NewSimpleStrictParameter(name, elem)
		}
		return NewSimpleParameter(name, elem)
	}
}
