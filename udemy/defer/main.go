package main

import "fmt"

func helloWorld() {
	defer fmt.Println("Hello")
	defer fmt.Println("Again. Last added in stack. First out in stack.")
	fmt.Println("World")
}

func main() {
	helloWorld()
}
