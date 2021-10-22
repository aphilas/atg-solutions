package sqrt

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	// float64(e) to prevent infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

const Threshold = 1e-10

func Sqrt(x float64) (float64, error) {
	y, z := 0.0, 1.0

	if x < 0 {
		e := ErrNegSqrt(x)
		return 0, e
	}

	for {
		z -= (z*z - x) / (2 * z)

		if math.Abs(y-z) < Threshold {
			return z, nil
		}

		y = z
	}
}
