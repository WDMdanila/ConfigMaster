package parameters

import (
	"fmt"
	"testing"
)

func TestJSONParameterReaderRead(t *testing.T) {
	expected := 11
	reader := NewJSONParameterReader("config.json")
	paramMap := reader.Read()
	for key, val := range paramMap {
		fmt.Printf("%v: %v\n", key, string(val.GetAsJSON()))
	}
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderReadFail1(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}

func TestJSONParameterReaderReadFail2(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config2.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}

func TestJSONParameterReaderReadFail3(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config3.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}

func TestJSONParameterReaderReadFail6(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config6.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}

func TestJSONParameterReaderReadFail7(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config7.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}

func TestJSONParameterReaderReadWork(t *testing.T) {
	expected := 1
	reader := NewJSONParameterReader("wrong_config4.json")
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderReadWork2(t *testing.T) {
	expected := 1
	reader := NewJSONParameterReader("wrong_config5.json")
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}
