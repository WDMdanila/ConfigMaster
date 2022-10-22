package parameters

import (
	"bytes"
	"testing"
)

func TestNamedParameterName(t *testing.T) {
	parameter := NamedParameter{name: "name"}
	if parameter.Name() != "name" {
		t.Fail()
	}
}

func TestSimpleParameterStringValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":"value"}`)
	parameter = NewSimpleParameter("value", "value")
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterStringValueSet(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":"value"}`)
	parameter = NewSimpleParameter("value", "qwe")
	parameter.Set(expected)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterIntValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":1}`)
	parameter = NewSimpleParameter("value", 1)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterIntValueSet(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":1}`)
	parameter = NewSimpleParameter("value", 0)
	parameter.Set(expected)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterBoolValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":true}`)
	parameter = NewSimpleParameter("value", true)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterBoolValueSet(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":true}`)
	parameter = NewSimpleParameter("value", false)
	parameter.Set(expected)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterJsonValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":{"field1":"value 1","field2":true,"field3":1}}`)
	parameter = NewSimpleParameter("value", []byte(`{"field1":"value 1","field2":true,"field3":1}`))
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterJsonValueSet(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":{"field1":"value 1","field2":true,"field3":1}}`)
	parameter = NewSimpleParameter("value", []byte(`{}`))
	parameter.Set(expected)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterFloatValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":3.141592653589793}`)
	parameter = NewSimpleParameter("value", float64(0))
	parameter.Set(expected)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
