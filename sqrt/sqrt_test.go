package sqrt_test

import (
	"math"
	"testing"

	"github.com/nevilleomangi/atg-solutions/sqrt"
)

func TestSqrt(t *testing.T) {
	t.Run("returns sqrt", func(t *testing.T) {
		cases := []struct {
			n    float64
			want float64
		}{
			{2, 1.41421356237},
			{16.0, 4.0},
			{57.0, 7.54983443527},
		}

		for _, c := range cases {
			got, err := sqrt.Sqrt(c.n)

			assertNoError(t, err)

			if !(math.Abs(got-c.want) < sqrt.Threshold) {
				t.Fatalf("got %g, want %g", got, c.want)
			}
		}
	})

	t.Run("returns error if argument < 0", func(t *testing.T) {
		_, err := sqrt.Sqrt(-1)
		want := sqrt.ErrNegSqrt(-1)
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
