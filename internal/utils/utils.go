package utils

import (
	"log"
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
		log.Panicf("could not read directory: %v, erorr: %v", dirPath, err)
	}
	return files
}

func GetFilenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
