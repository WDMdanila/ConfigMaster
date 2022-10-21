package utils

import (
	"bytes"
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

func TestGetAsJSON(t *testing.T) {
	expected := []byte(`{"test":"test"}`)
	res := GetAsJSON("test", "test")
	if !bytes.Equal(res, expected) {
		t.Fatalf("encoding data as JSON: expected: %v, got: %v", string(expected), string(res))
	}
}

func TestGetAsJSONError(t *testing.T) {
	defer func() { _ = recover() }()
	GetAsJSON("test", make(chan int))
	t.Fail()
}

func TestDecodeJSON(t *testing.T) {
	_ = DecodeJSON[interface{}]([]byte(`{"test":"test"}`))
}

func TestDecodeJSONFail(t *testing.T) {
	defer func() { _ = recover() }()
	DecodeJSON[[]byte](nil)
	t.Fail()
}
