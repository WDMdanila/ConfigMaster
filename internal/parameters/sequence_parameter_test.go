package parameters

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestArithmeticSequenceParameterValue(t *testing.T) {
	var parameter Parameter
	expected := float64(0)
	parameter = NewArithmeticSequenceParameter("value", 0, 1)
	res := parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	expected = 1
	res = parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestArithmeticSequenceParameterDescribe(t *testing.T) {
	expected := map[string]interface{}{"value": map[string]interface{}{"value": float64(1), "increment": float64(10), "parameter_type": "arithmetic sequence"}}
	var parameter Parameter
	parameter = NewArithmeticSequenceParameter("value", 1, 10)
	res := parameter.Describe()
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected: %v, got: %v", expected, res)
	}
}

func TestArithmeticSequenceParameterSet(t *testing.T) {
	var parameter Parameter
	parameter = NewArithmeticSequenceParameter("value", 1, 10)
	err := parameter.Set(map[string]float64{"increment": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(map[string]float64{"increment": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(nil)
	if err == nil {
		t.Fatal(err)
	}
}

func TestArithmeticSequenceParameterSet2(t *testing.T) {
	var parameter Parameter
	parameter = NewArithmeticSequenceParameter("value", 1, 10)
	err := parameter.Set(map[string]float64{"increment": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(map[string]interface{}{"increment": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(nil)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGeometricSequenceParameterValue(t *testing.T) {
	var parameter Parameter
	expected := float64(1)
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	res := parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	expected = 10
	res = parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestGeometricSequenceParameterDescribe(t *testing.T) {
	expected := map[string]interface{}{"value": map[string]interface{}{"value": float64(1), "multiplier": float64(10), "parameter_type": "geometric sequence"}}
	var parameter Parameter
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	res := parameter.Describe()
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected: %v, got: %v", expected, res)
	}
}

func TestGeometricSequenceParameterSet(t *testing.T) {
	var parameter Parameter
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	err := parameter.Set(map[string]float64{"multiplier": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(map[string]float64{"multiplier": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(nil)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGeometricSequenceParameterSet2(t *testing.T) {
	var parameter Parameter
	parameter = NewGeometricSequenceParameter("value", 1, 10)
	err := parameter.Set(map[string]float64{"multiplier": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(map[string]interface{}{"multiplier": 1, "value": 1})
	if err != nil {
		t.Fatal(err)
	}
	err = parameter.Set(nil)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGeometricFloatSequenceParameterValue(t *testing.T) {
	var parameter Parameter
	expected := float64(10)
	parameter = NewGeometricSequenceParameter("value", 10, 0.1)
	res := parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	expected = 1
	res = parameter.Value()
	if res.(float64) != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestRandomParameterValue(t *testing.T) {
	var parameter Parameter
	expected := 0
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 1)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestRandomParameterDescribe(t *testing.T) {
	expected := map[string]interface{}{"value": map[string]interface{}{"min": 1, "max": 10, "parameter_type": "random"}}
	var parameter Parameter
	parameter = NewRandomParameter("value", 1, 10)
	res := parameter.Describe()
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected: %v, got: %v", expected, res)
	}
}

func TestRandomParameterSet(t *testing.T) {
	var parameter Parameter
	expected := 84
	rand.Seed(10)
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set(map[string]float64{"min": 10, "max": 100})
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestRandomParameterSetFail(t *testing.T) {
	var parameter Parameter
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set(map[string]interface{}{"min": "fail", "max": 100})
	if err == nil {
		t.Fatal(err)
	}
}

func TestRandomParameterSetFail2(t *testing.T) {
	var parameter Parameter
	parameter = NewRandomParameter("value", 0, 1)
	err := parameter.Set(map[string]interface{}{"min": 1, "max": "100"})
	if err == nil {
		t.Fatal(err)
	}
}

func TestRandomParameterRandomValue(t *testing.T) {
	var parameter Parameter
	expected := 4
	rand.Seed(0)
	parameter = NewRandomParameter("value", 0, 10)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	rand.Seed(1)
	expected = 1
	res = parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}
