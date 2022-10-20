package parameters

import (
	"testing"
)

func TestJSONParameterReaderRead(t *testing.T) {
	expected := 11
	reader := NewJSONParameterReader("config.json")
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderUnknownType(t *testing.T) {
	reader := NewJSONParameterReader("wrong_config.json")
	defer func() { _ = recover() }()
	reader.Read()
	t.Fail()
}
