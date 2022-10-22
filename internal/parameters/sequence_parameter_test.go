package parameters

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestArithmeticSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":0}`)
	parameter = NewArithmeticSequenceParameter("value", 0, 1)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res = parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestArithmeticFloatSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":0}`)
	parameter = NewArithmeticSequenceParameter("value", 0, 0.1)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":0.1}`)
	res = parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestGeometricSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":1}`)
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":10}`)
	res = parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestGeometricFloatSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":10}`)
	parameter = NewGeometricSequenceParameter("value", 10, 0.1)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res = parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":0}`)
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 1)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterRandomAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":4}`)
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 10)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	rand.Seed(1)
	expected = []byte(`{"value":1}`)
	res = parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
