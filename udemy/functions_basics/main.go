package main

import "fmt"

func aFunction() {
	fmt.Println("This is a defined function!")
}

func returnMyString(s string) string {
	return "function returned " + s
}

func main() {
	aFunction()
	fmt.Println(returnMyString("ye"))
	func() {
		fmt.Println("This is an anonymous function!")
	}()
}
