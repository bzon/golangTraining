package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// x, y = computeX(y), computeY(x, y)
		x, y = y, x+y
		fmt.Println(x, y)
	}
	return x
}

// func computeX(y int) int {
// 	fmt.Println("x is now", y)
// 	return y
// }

// func computeY(x, y int) int {
// 	newY := x + y
// 	fmt.Println("y is now", newY)
// 	return newY
// }
