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
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res, err = parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestArithmeticFloatSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":0}`)
	parameter = NewArithmeticSequenceParameter("value", 0, 0.1)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":0.1}`)
	res, err = parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestArithmeticSequenceParameterSet(t *testing.T) {
	var parameter Parameter
	parameter = NewArithmeticSequenceParameter("value", 1, 10)
	err := parameter.Set([]byte(`{"increment":1,"value":1}`))
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set([]byte(`{"increment":1}`))
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set([]byte(`{}`))
	if err == nil {
		t.Fatal(err)
	}
}

func TestGeometricSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":1}`)
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":10}`)
	res, err = parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestGeometricSequenceParameterSet(t *testing.T) {
	var parameter Parameter
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	err := parameter.Set([]byte(`{"multiplier":1,"value":1}`))
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set([]byte(`{"multiplier":1}`))
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set([]byte(`{}`))
	if err == nil {
		t.Fatal(err)
	}
}

func TestGeometricFloatSequenceParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":10}`)
	parameter = NewGeometricSequenceParameter("value", 10, 0.1)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res, err = parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":0}`)
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 1)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterSet(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":84}`)
	rand.Seed(10)
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set([]byte(`{"min":10,"max":100}`))
	if err != nil {
		t.Fatal(err)
	}
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterSetFail(t *testing.T) {
	var parameter Parameter
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set([]byte(`{"min":"fail","max":100}`))
	if err == nil {
		t.Fatal(err)
	}
}

func TestRandomParameterSetFail2(t *testing.T) {
	var parameter Parameter
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set([]byte(`{"min":1,"max":"100"}`))
	if err == nil {
		t.Fatal(err)
	}
}

func TestRandomParameterRandomAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":4}`)
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 10)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	rand.Seed(1)
	expected = []byte(`{"value":1}`)
	res, err = parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
