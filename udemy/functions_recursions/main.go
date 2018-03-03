package main

import (
	"fmt"
)

func factorial(n int) int {
	fmt.Printf("n is now %d\n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("factorial of 5 is", factorial(5))
}

// formula: factorial of 5 is, 5 * 4 * 3 * 2 * 1
