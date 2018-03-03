package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// value receiver
func (p person) speak() {
	fmt.Println("My name is", p.first)
}

// pointer receiver
func (p *person) changeName(newName string) {
	p.first = newName
}

type android struct {
	person
	version string
}

func main() {
	p1 := person{"john", "sazon", 28}
	fmt.Println(p1.first, p1.last, p1.age)
	p1.speak()
	p1.changeName("bryan")
	fmt.Println(p1.first, p1.last, p1.age)

	p2 := android{
		person: person{
			first: "john",
			last:  "wick",
			age:   999,
		},
		version: "v1.0",
	}
	fmt.Println(p2.first, p2.person.last)
}
