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

func TestSimpleParameterJSONValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":{"field1":"value 1","field2":true,"field3":1}}`)
	parameter = NewSimpleParameter("value", []byte(`{"field1":"value 1","field2":true,"field3":1}`))
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleParameterJSONValueSet(t *testing.T) {
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

func TestSimpleStrictParameterJSONValueSetFail(t *testing.T) {
	var parameter Parameter
	parameter = NewSimpleStrictParameter("value", []byte(`{}`))
	err := parameter.Set([]byte(`{"value": 1}`))
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSimpleStrictParameterArrayValueGetAsJSON(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":[1,2,3]}`)
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewSimpleStrictParameter("value", values)
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleStrictParameterArrayValueSetFail(t *testing.T) {
	var parameter Parameter
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewSimpleStrictParameter("value", values)
	err := parameter.Set([]byte(`{"value": 1}`))
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSimpleStrictParameterJSONValueArray(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":[1,2,3]}`)
	parameter = NewSimpleStrictParameter("value", []byte(`[1,2,3]`))
	res := parameter.GetAsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSimpleStrictParameterJSONValueSetArray(t *testing.T) {
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
