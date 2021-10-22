package fib_test

import (
	"testing"

	"github.com/nevilleomangi/atg-solutions/fib"
)

func TestWalker(t *testing.T) {
	want := [...]int{0, 1, 1, 2, 3, 5}

	f := fib.Fibonacci()
	for i, n := range want {
		got := f()
		if n != got {
			t.Fatalf("i = %d, got %d, want %d", i, got, n)
		}
	}
}
