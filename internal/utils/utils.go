package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFilesWithExtInDirectory(dirPath string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(info.Name()) == "."+ext {
			files = append(files, path)
		}
		return err
	})
	return files, err
}

func GetFilenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetAsJSON[T any](key string, value T) ([]byte, error) {
	tmp := map[string]T{key: value}
	return json.Marshal(tmp)
}

func DecodeJSON[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}

func ExtractFromJSON[T any](data []byte, field string) (T, error) {
	var res T
	tmp, err := DecodeJSON[map[string]interface{}](data)
	if err != nil {
		return res, err
	}
	switch value := tmp[field].(type) {
	case T:
		return value, nil
	default:
		return res, fmt.Errorf("could not parse %v, got type %T", value, value)
	}
}

func ExtractFileNameAndPath(fileName string) (string, string) {
	fileName = strings.ReplaceAll(fileName, `\`, "/")
	fileName = strings.ReplaceAll(fileName, "//", "/")
	log.Printf("found config: %v\n", fileName)
	folderNameIndex := strings.Index(fileName, "/")
	return fileName, GetFilenameWithoutExt(fileName[folderNameIndex+1:])
}
