package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(2)

	for {
		t := z
		z = z - (z*z-x)/(2*z)
		if diff := t - z; diff < 0.000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
