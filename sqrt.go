package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	v, z := 0., 1.
	for {
		z, v = z - (z * z - x) / (2 * z), z
		if math.Abs(z-v) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}