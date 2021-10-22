package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	got := Sqrt(16)
	want := 4.0

	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
