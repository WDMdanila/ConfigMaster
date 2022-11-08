package parameters

import (
	"bytes"
	"encoding/json"
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

func (r *JSONParameterReader) Read() map[string]Parameter {
	data := parseJSONFile(r.filePath)
	return r.fromParsedJSON(data)
}

func (r *JSONParameterReader) fromParsedJSON(data map[string]interface{}) map[string]Parameter {
	res := map[string]Parameter{}
	for paramName, paramData := range data {
		parameter := parseParameter(paramName, paramData, r.strictTypes)
		res[parameter.Name()] = parameter
	}
	return res
}

func NewJSONParameterReader(filePath string, strictTypes bool) ParameterReader {
	return &JSONParameterReader{filePath: filePath, strictTypes: strictTypes}
}

func parseJSONFile(filePath string) map[string]interface{} {
	rawData := getCleanedRawData(filePath)
	var data map[string]interface{}
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic(err)
	}
	return data
}

func getCleanedRawData(filePath string) []byte {
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}
	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))
	return byteValue
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
