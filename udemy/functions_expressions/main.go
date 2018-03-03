package main

import (
	"fmt"
)

func main() {
	helloWorld := func() {
		fmt.Println("Hello World")
	}
	fmt.Printf("%T \n", helloWorld)
	helloWorld()
	fmt.Println()

	// invoke directly ..
	fmt.Printf("%T \n", makeGreetings)
	fmt.Println(makeGreetings()())
	// or using func expressions
	greetings := makeGreetings()
	fmt.Println(greetings())
}

// with closure style
func makeGreetings() func() string {
	return func() string {
		return "Greetings!"
	}
}
