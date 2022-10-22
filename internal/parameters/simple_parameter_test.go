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
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
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
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
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
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
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
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleStrictParameterJsonValueSetFail(t *testing.T) {
	var parameter Parameter
	parameter = NewSimpleStrictParameter("value", []byte(`{}`))
	err := parameter.Set([]byte(`{"value": 1}`))
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSimpleStrictParameterJsonValueArray(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":[1,2,3]}`)
	parameter = NewSimpleStrictParameter("value", []byte(`[1,2,3]`))
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleStrictParameterJsonValueSetArray(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":[1,2,3]}`)
	parameter = NewSimpleStrictParameter("value", []byte(`[1,2]`))
	err := parameter.Set([]byte(`{"value":[1,2,3]}`))
	if err != nil {
		t.Fatal()
	}
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterFloatValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":3.141592653589793}`)
	parameter = NewSimpleParameter("value", float64(0))
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
