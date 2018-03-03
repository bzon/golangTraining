package main

import "fmt"

func main() {
	fmt.Println(sum(5, 3, 2))
	printing("hello", 55)
}

func sum(n ...float64) float64 {
	fmt.Printf("n type is %T\n", n)
	var s float64
	for _, v := range n {
		s += v
	}
	return s
}

func printing(i ...interface{}) {
	fmt.Printf("s type is %T %v\n", i, i)
}
