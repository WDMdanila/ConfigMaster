package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFilesWithExtInDirectory(dirPath string, ext string) []string {
	var files []string
	directoryFiles, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("could not read directory: %v, erorr: %v", dirPath, err)
	}
	for _, file := range directoryFiles {
		if !file.IsDir() && filepath.Ext(file.Name()) == "."+ext {
			files = append(files, file.Name())
		}
	}
	return files
}

func GetFilenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
