package main

import (
	"errors"
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	// Handle negative input
	if x < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}

	// Newton's method
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
