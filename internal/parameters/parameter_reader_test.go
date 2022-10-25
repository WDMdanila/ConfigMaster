package parameters

import (
	"testing"
)

func TestJSONParameterReaderReadNonExistent(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/does_not_exist", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderRead(t *testing.T) {
	expected := 11
	reader := NewJSONParameterReader("test_configs/correct_config1.json", true)
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderReadNonStrict(t *testing.T) {
	expected := 11
	reader := NewJSONParameterReader("test_configs/correct_config1.json", false)
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderReadFail1(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadFail2(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config2.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadFail3(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config3.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadFail4(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config4.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadFail5(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config5.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadFail6(t *testing.T) {
	reader := NewJSONParameterReader("test_configs/wrong_config6.json", true)
	defer func() { _ = recover() }()
	reader.Read()
	t.Fatal()
}

func TestJSONParameterReaderReadWork(t *testing.T) {
	expected := 1
	reader := NewJSONParameterReader("test_configs/correct_config3.json", true)
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}

func TestJSONParameterReaderReadWork2(t *testing.T) {
	expected := 1
	reader := NewJSONParameterReader("test_configs/correct_config2.json", true)
	paramMap := reader.Read()
	if len(paramMap) != expected {
		t.Fatalf("should be %v parameters, only %v present", len(paramMap), expected)
	}
}
