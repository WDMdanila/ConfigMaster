package utils

import (
	"testing"
)

func TestFindFilesWithExtInDirectory(t *testing.T) {
	expected := 2
	modFile := FindFilesWithExtInDirectory("./", "txt")
	modCount := len(modFile)
	if modCount != expected {
		t.Fatalf("found wrong number of .txt files in directory: %v, expected: %v", modCount, expected)
	}
}

func TestFindFilesWithExtInNonExistentDirectory(t *testing.T) {
	defer func() { _ = recover() }()
	FindFilesWithExtInDirectory("./does_not_exist", "txt")
	t.Fail()
}

func TestGetFilenameWithoutExt(t *testing.T) {
	expected := "test"
	filename := expected + ".tst"
	res := GetFilenameWithoutExt(filename)
	if res != expected {
		t.Fatalf("filename without extension is wrong, expected: %v, got: %v", expected, res)
	}
}
