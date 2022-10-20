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
	parameter = NewRandomSelectionParameter([]int{1, 2, 3})
	res := parameter.ToJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestSequentialSelectionParameterAsJSON(t *testing.T) {
	expected := []byte(`{"value":1}`)
	var parameter Parameter
	parameter = NewSequentialSelectionParameter([]int{1, 2})
	res := parameter.ToJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":2}`)
	res = parameter.ToJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	expected = []byte(`{"value":1}`)
	res = parameter.ToJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
