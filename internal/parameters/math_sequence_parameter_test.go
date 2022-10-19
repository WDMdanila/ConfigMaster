package parameters

import (
	"bytes"
	"testing"
)

func TestArithmeticSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":0}`)
	parameter = NewArithmeticSequenceParameter(0, 1)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res = parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestArithmeticFloatSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":0}`)
	parameter = NewArithmeticSequenceParameter[float64](0, 0.1)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":0.1}`)
	res = parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestGeometricSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":1}`)
	parameter = NewGeometricSequenceParameter(1, 10)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":10}`)
	res = parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestGeometricFloatSequenceParameterAsJSON(t *testing.T) {

	var parameter Parameter
	var expected = []byte(`{"value":10}`)
	parameter = NewGeometricSequenceParameter[float64](10, 0.1)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res = parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
