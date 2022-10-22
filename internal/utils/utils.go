package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFilesWithExtInDirectory(dirPath string, ext string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(info.Name()) == "."+ext {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		panic(fmt.Errorf("could not read directory: %v, erorr: %v", dirPath, err))
	}
	return files
}

func GetFilenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetAsJSON[T any](key string, value T) []byte {
	tmp := map[string]T{key: value}
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func DecodeJSON[T any](data []byte) T {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func ExtractFromJSON[T any](data []byte, field string) (T, error) {
	var res T
	tmp := DecodeJSON[map[string]interface{}](data)
	switch value := tmp[field].(type) {
	case T:
		return value, nil
	default:
		return res, fmt.Errorf("could not parse %v, got type %T", value, value)
	}
}
