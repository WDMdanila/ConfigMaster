package parameters

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestRandomSelectionParameterAsJSON(t *testing.T) {
	expected := []byte(`{"value":1}`)
	var parameter Parameter
	rand.Seed(0)
	values := make([]interface{}, 0)
	values = append(values, 1, 2, 3)
	parameter = NewRandomSelectionParameter("value", values)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSequentialSelectionParameterAsJSON(t *testing.T) {
	expected := []byte(`{"value":1}`)
	var parameter Parameter
	values := make([]interface{}, 0)
	values = append(values, 1, 2)
	parameter = NewSequentialSelectionParameter("value", values)
	res, err := parameter.GetAsJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":2}`)
	res, err = parameter.GetAsJSON()
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

func TestRandomSelectionParameterSet(t *testing.T) {
	expected := []byte(`{"value":4}`)
	var parameter Parameter
	rand.Seed(0)
	values := make([]interface{}, 0)
	values = append(values, 1)
	parameter = NewRandomSelectionParameter("value", values)
	err := parameter.Set([]byte(`{"values": [4]}`))
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
