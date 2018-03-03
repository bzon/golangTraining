package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := 0
	fmt.Println(xPtr)
	one(&xPtr)
	fmt.Println(xPtr) // x is 1
}
