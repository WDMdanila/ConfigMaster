package parameters

import (
	"bytes"
	"testing"
)

func TestSimpleParameterStringValueAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":"value"}`)
	parameter = &SimpleParameter[string]{"value"}
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterIntValueAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":1}`)
	parameter = &SimpleParameter[int]{1}
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterBoolValueAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":true}`)
	parameter = &SimpleParameter[bool]{true}
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterJsonValueAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":{"field1":"value 1","field2":true,"field3":1}}`)
	parameter = NewJSONParameter([]byte(`{"field1":"value 1","field2":true,"field3":1}`))
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterFloatValueAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":3.141592653589793}`)
	parameter = &SimpleParameter[float64]{3.141592653589793}
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
