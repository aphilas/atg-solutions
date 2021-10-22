package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	t.Run("returns sqrt", func(t *testing.T) {
		got, err := Sqrt(16)
		want := 4.0
		assertNoError(t, err)

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})

	t.Run("returns error if argument < 0", func(t *testing.T) {
		_, err := Sqrt(-1)
		want := ErrNegSqrt(-1)
		assertError(t, err, want)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got error, wanted nil")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("got nil, wanted error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
