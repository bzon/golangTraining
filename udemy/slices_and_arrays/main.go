package main

import "fmt"

func main() {
	// array
	a := [5]int{1, 2, 3, 4, 5}
	// a = append(a, 6) // error! array length is fixed
	fmt.Println(a)
	// slice is like a reference to array
	s := []int{}
	fmt.Println(s)
	// it can be resized!
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	s2 := make([]int, 5)
	fmt.Println(len(s2))
	s2 = append(s2, 1)
	fmt.Println(s2)
	fmt.Println(len(s2)) // it increased

	// Strings are a slice of bytes or uint8
	fmt.Printf("%T %d\n", "mySlice"[2], "mySlice"[2])
	fmt.Printf("%T %s\n", string("mySlice"[2]), string("mySlice"[2]))

	// Different type of declaration or creation of a slice
	s3 := []int{}
	fmt.Println(s3)
	s4 := make([]int, 1, 5)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	var s5 []int
	fmt.Println(s5)

}
