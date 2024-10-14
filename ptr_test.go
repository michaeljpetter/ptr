package ptr_test

import (
	"github.com/michaeljpetter/ptr"
	"testing"
)

func TestTo(t *testing.T) {
	in := "a value"

	out := ptr.To(in)

	if out == nil {
		t.Errorf("wrong value %v, expected -> %v", out, in)
	}
	if *out != in {
		t.Errorf("wrong value -> %v, expected -> %v", *out, in)
	}
}

func TestOrZero(t *testing.T) {
	in := "a value"

	out := ptr.OrZero(&in)

	if out == nil {
		t.Errorf("wrong value %v, expected -> %v", out, in)
	}
	if *out != in {
		t.Errorf("wrong value -> %v, expected -> %v", *out, in)
	}
}

func TestOrZeroNil(t *testing.T) {
	var in *int

	out := ptr.OrZero(in)

	exp := 0

	if out == nil {
		t.Errorf("wrong value %v, expected -> %v", out, exp)
	}
	if *out != exp {
		t.Errorf("wrong value -> %v, expected -> %v", *out, exp)
	}
}
