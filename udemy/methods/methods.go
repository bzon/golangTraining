package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) sum() float64 {
	return v.X + v.Y
}

func (v *Vertex) sumPointer() float64 {
	return v.X + v.Y
}

type MyFloat float64

func (f MyFloat) plus(x MyFloat) MyFloat {
	return f + x
}

func main() {

	// using literal declaration
	fmt.Println(Vertex{3, 4})
	fmt.Println(Vertex{3, 4}.sum())

	// using pointer receivers with struct type
	var a = Vertex{5, 4}
	fmt.Println(a.sumPointer())

	// using value receivers with float64 type
	var i MyFloat = 1.5
	fmt.Println(i.plus(5))
}