package main

import (
	"fmt"
	"math"
)

const threshold = 1e-10

func Sqrt(x float64) float64 {
	y, z := 0.0, 1.0

	for {
		z -= (z*z - x) / (2 * z)

		if math.Abs(y-z) < threshold {
			return z
		}

		y = z
	}
}

func main() {
	fmt.Println(Sqrt(16))
}
