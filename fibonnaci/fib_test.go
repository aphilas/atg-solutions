package main

import (
	"testing"
)

func TestWalker(t *testing.T) {
	want := [...]int{0, 1, 1, 2, 3, 5}

	f := Fibonacci()
	for i, n := range want {
		got := f()
		if n != got {
			t.Fatalf("i = %d, got %d, want %d", i, got, n)
		}
	}
}
