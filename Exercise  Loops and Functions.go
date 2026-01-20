package main

import (
	"fmt"
	"math"
)

// Sqrt calculates square root using Newton's method
func Sqrt(x float64) float64 {
	z := 1.0 // initial guess

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println("Iteration", i+1, ":", z)
	}

	return z
}

func main() {
	x := 2.0
	mySqrt := Sqrt(2.0)

	fmt.Println("My Sqrt :", mySqrt)
	fmt.Println("math.Sqrt:", math.Sqrt(x))
}
