package utils

import (
	"testing"
)

func TestFindFilesWithExtInDirectory(t *testing.T) {
	modFile := FindFilesWithExtInDirectory("./", "txt")
	modCount := len(modFile)
	if modCount != 1 {
		t.Fatalf("found multiple .txt files in directory: %v", modCount)
	}
}

func TestGetFilenameWithoutExt(t *testing.T) {
	var expected = "test"
	var filename = expected + ".tst"
	res := GetFilenameWithoutExt(filename)
	if res != expected {
		t.Fatalf("filename without extension is wrong, expected: %v, got: %v", expected, res)
	}
}
