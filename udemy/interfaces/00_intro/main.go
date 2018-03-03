package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type rectangle struct {
	sides float64
}

func main() {
	s := square{10}
	// using the method directly
	fmt.Println(s.area())
	fmt.Printf("%T", s)
	// square is a shape, and func info can only take type that implements shape
	info(s)
	c := circle{5}
	// using the method directly
	fmt.Println(c.area())
	// circle is a shape, and func info can only take type that implements shape
	info(c)
	// Err
	// info(rectangle{5})
}
