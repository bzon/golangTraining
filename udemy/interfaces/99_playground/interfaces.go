package main

import "fmt"

type Word string

type MyFloat float64

type Person struct {
	Name string
	Age  int
}

// printType is an implicitly defined interface!
func (w Word) printType() {
	fmt.Printf("%s type is %T\n", w, w)
}

func (f MyFloat) printType() {
	fmt.Printf("%f type is %T\n", f, f)
}

func (p Person) printType() {
	fmt.Printf("%s with age %d type is %T\n", p.Name, p.Age, p)
}

// MyInterface implementation
type MyInterface interface {
	printTypeAgain()
}

func (w Word) printTypeAgain() {
	fmt.Printf("%s type is %T\n", w, w)
}

func (f MyFloat) printTypeAgain() {
	fmt.Printf("%f type is %T\n", f, f)
}

func (p Person) printTypeAgain() {
	fmt.Printf("%s with age %d type is %T\n", p.Name, p.Age, p)
}

func main() {
	Word("hello").printType()
	MyFloat(5).printType()
	Person{"Bryan", 28}.printType()
	var i MyInterface = Word("hello again")
	i.printTypeAgain()
}
