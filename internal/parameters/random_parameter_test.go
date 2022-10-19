package parameters

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestRandomParameterAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":0}`)
	rand.Seed(0)
	parameter = NewRandomParameter(0, 1)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}

func TestRandomParameterRandomAsJSON(t *testing.T) {
	var parameter Parameter
	var expected = []byte(`{"value":4}`)
	rand.Seed(0)
	parameter = NewRandomParameter(0, 10)
	res := parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
	rand.Seed(1)
	expected = []byte(`{"value":1}`)
	res = parameter.AsJSON()
	if !bytes.Equal(res, expected) {
		t.Fatalf("parameter json %v does not equal to %v", string(res), string(expected))
	}
}
