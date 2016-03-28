package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.4
	// z := float64(1)
	var r float64
	for i := 0; i < 1000; i++ {
		r = z - ((z * z) - x)/(2 * z)
		if r - z < 0.000001 && r - z > -0.1 {
			return r
		}
		z = r
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}