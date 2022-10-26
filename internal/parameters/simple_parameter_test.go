package parameters

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNamedParameterName(t *testing.T) {
	parameter := NamedParameter{name: "name"}
	if parameter.Name() != "name" {
		t.Fatal()
	}
}

func TestSimpleParameterStringValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := "value"
	parameter = NewSimpleParameter("value", "value")
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterStringValueDescribe(t *testing.T) {
	var parameter Parameter
	expected := map[string]string{"value": "value"}
	parameter = NewSimpleParameter("value", "value")
	res := parameter.Describe()
	for key, val := range res {
		if val.(string) != expected[key] {
			t.Fatalf("parameter value %v does not equal to %v", res, expected)
		}
	}
}

func TestSimpleParameterStringValueSet(t *testing.T) {
	var parameter Parameter
	expected := "new value"
	parameter = NewSimpleParameter("value", "qwe")
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterIntValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := 1
	parameter = NewSimpleParameter("value", 1)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterIntValueSet(t *testing.T) {
	var parameter Parameter
	expected := 1
	parameter = NewSimpleParameter("value", 0)
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterBoolValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := true
	parameter = NewSimpleParameter("value", true)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterBoolValueSet(t *testing.T) {
	var parameter Parameter
	expected := true
	parameter = NewSimpleParameter("value", false)
	err := parameter.Set(true)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterJSONValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"field1":"value 1","field2":true,"field3":1}`)
	parameter = NewSimpleParameter("value", []byte(`{"field1":"value 1","field2":true,"field3":1}`))
	res := parameter.Value()
	if !bytes.Equal(res.([]byte), expected) {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
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
	res := parameter.Value()
	if !bytes.Equal(res.([]byte), expected) {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleStrictParameterJSONValueSetFail(t *testing.T) {
	var parameter Parameter
	parameter = NewSimpleStrictParameter("value", []byte(`{}`))
	err := parameter.Set(1)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSimpleStrictParameterArrayValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := []int{1, 2, 3}
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewSimpleStrictParameter("value", values)
	res := parameter.Value()
	for i, value := range res.([]interface{}) {
		switch val := value.(type) {
		case int:
			if val != expected[i] {
				t.Fatalf("parameter value %v does not equal to %v", val, expected[i])
			}
		default:
			t.Fatal()
		}
	}
}

func TestSimpleStrictParameterArrayValueSetFail(t *testing.T) {
	var parameter Parameter
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewSimpleStrictParameter("value", values)
	err := parameter.Set(1)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSimpleStrictParameterJSONValueArray(t *testing.T) {
	var parameter Parameter
	var expected interface{}
	expected = []byte(`[1,2,3]`)
	parameter = NewSimpleStrictParameter("value", []byte(`[1,2,3]`))
	res := parameter.Value()
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleStrictParameterJSONValueSetArray(t *testing.T) {
	var parameter Parameter
	expected := []byte(`{"value":[1,2,3]}`)
	parameter = NewSimpleStrictParameter("value", []byte(`[1,2]`))
	err := parameter.Set([]byte(`[1,2,3]`))
	if err != nil {
		t.Fatal()
	}
	res := parameter.Value()
	if bytes.Equal(res.([]byte), expected) {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}

func TestSimpleParameterFloatValueGetValue(t *testing.T) {
	var parameter Parameter
	expected := 3.141592653589793
	parameter = NewSimpleParameter("value", float64(0))
	err := parameter.Set(expected)
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter value %v does not equal to %v", res, expected)
	}
}
