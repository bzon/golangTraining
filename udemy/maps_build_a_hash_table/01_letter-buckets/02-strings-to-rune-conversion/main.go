package main

import "fmt"

func main() {
	letter := rune("A"[0]) // rune is an alias for int32
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
