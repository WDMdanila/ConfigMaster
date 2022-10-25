package utils

import (
	"bytes"
	"testing"
)

func TestFindFilesWithExtInDirectory(t *testing.T) {
	expected := 2
	modFile, err := FindFilesWithExtInDirectory("./", "txt")
	if err != nil {
		t.Fatal(err)
	}
	modCount := len(modFile)
	if modCount != expected {
		t.Fatalf("found wrong number of .txt files in directory: %v, expected: %v", modCount, expected)
	}
}

func TestFindFilesWithExtInNonExistentDirectory(t *testing.T) {
	_, err := FindFilesWithExtInDirectory("./does_not_exist", "txt")
	if err == nil {
		t.Fatal("expected error")
	}
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
	res, err := GetAsJSON("test", "test")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("encoding data as JSON: expected: %v, got: %v", string(expected), string(res))
	}
}

func TestGetAsJSONError(t *testing.T) {
	_, err := GetAsJSON("test", make(chan int))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDecodeJSON(t *testing.T) {
	_, err := DecodeJSON[interface{}]([]byte(`{"value":"test"}`))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeJSONFail(t *testing.T) {
	_, err := DecodeJSON[[]byte](nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestExtractFromJSON(t *testing.T) {
	expected := float64(1)
	res, err := ExtractFromJSON[float64]([]byte(`{"value":1}`), "value")
	if err == nil && res != expected {
		t.Fatal()
	}
}

func TestExtractFromJSONFail(t *testing.T) {
	_, err := ExtractFromJSON[string]([]byte(`{"value":1}`), "value")
	if err == nil {
		t.Fatal()
	}
}

func TestExtractFromJSONFail2(t *testing.T) {
	_, err := ExtractFromJSON[string]([]byte(`{`), "")
	if err == nil {
		t.Fatal()
	}
}
