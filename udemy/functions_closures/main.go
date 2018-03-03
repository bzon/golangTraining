package main

import (
	"fmt"
)

// incrementing using package scope
var x1 int

func increment1() int {
	x1++
	return x1
}

// incrementing by limiting the scope of the variable
func increment3() func() int {
	x3 := 0
	return func() int {
		x3++
		return x3
	}
}

func main() {
	fmt.Println("incrementing a var using package scope")
	fmt.Println(increment1())
	fmt.Println(increment1())

	fmt.Println("incrementing a var using a closure defined inside main")
	x2 := 0
	increment2 := func() int {
		x2++
		return x2
	}
	fmt.Println(increment2())
	fmt.Println(increment2())

	fmt.Println("incrementing a var using a closure using a function defined outside main")
	increment := increment3()
	fmt.Println(increment())
	fmt.Println(increment())
}

// closures: function is bound by a variable
