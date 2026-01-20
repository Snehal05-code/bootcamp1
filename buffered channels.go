package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer size = 2

	ch <- 1
	ch <- 2

	fmt.Println("Sent two values")

	// ❗ This third send overfills the buffer
	ch <- 3

	fmt.Println("This line will NOT execute")
}
