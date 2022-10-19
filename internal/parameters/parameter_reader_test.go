package parameters

import (
	"testing"
)

func TestJSONParameterReaderRead(t *testing.T) {
	var expected = 11
	reader := NewJSONParameterReader("config.json")
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}
