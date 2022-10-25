package parameters

import (
	"math/rand"
	"testing"
)

func TestRandomSelectionParameterValue(t *testing.T) {
	expected := 1
	var parameter Parameter
	rand.Seed(0)
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewRandomSelectionParameter("value", values)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestSequentialSelectionParameterValue(t *testing.T) {
	expected := 1
	var parameter Parameter
	values := make([]interface{}, 0)
	values = append(values, 1, 2)
	parameter = NewSequentialSelectionParameter("value", values)
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	expected = 2
	res = parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
	expected = 1
	res = parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestRandomSelectionParameterSet(t *testing.T) {
	expected := 4
	var parameter Parameter
	rand.Seed(0)
	values := make([]interface{}, 0)
	values = append(values, 1)
	parameter = NewRandomSelectionParameter("value", values)
	err := parameter.Set([]interface{}{4})
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestSequentialSelectionParameterSet(t *testing.T) {
	expected := 4
	var parameter Parameter
	rand.Seed(0)
	values := make([]interface{}, 0)
	values = append(values, 1)
	parameter = NewSequentialSelectionParameter("value", values)
	err := parameter.Set([]interface{}{4})
	if err != nil {
		t.Fatal(err)
	}
	res := parameter.Value()
	if res != expected {
		t.Fatalf("parameter json %v does not equal to %v", res, expected)
	}
}

func TestRandomSelectionParameterSetFail(t *testing.T) {
	var parameter Parameter
	values := make([]interface{}, 0)
	values = append(values, 1)
	parameter = NewRandomSelectionParameter("value", values)
	err := parameter.Set([]byte(`{"values": 4}`))
	if err == nil {
		t.Fatal("expected an error")
	}
}
