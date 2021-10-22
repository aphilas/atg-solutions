package main

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	// float64(e) to prevent infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

const threshold = 1e-10

func Sqrt(x float64) (float64, error) {
	y, z := 0.0, 1.0

	if x < 0 {
		e := ErrNegSqrt(x)
		return 0, e
	}

	for {
		z -= (z*z - x) / (2 * z)

		if math.Abs(y-z) < threshold {
			return z, nil
		}

		y = z
	}
}

func main() {
	n := 2.0
	s, err := Sqrt(n)
	fmt.Printf("Sqrt(%g) = %g\n", n, s)

	if err != nil {
		// fmt.Printf("Error generating Sqrt(%g)\n", n)
		fmt.Println(err)
	}

	n = -2.0
	_, err = Sqrt(n)

	if err != nil {
		fmt.Println(err)
	}
}
