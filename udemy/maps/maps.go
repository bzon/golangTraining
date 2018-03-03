package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m["a"])

	// Print all values
	for key, value := range m {
		fmt.Printf("Key %s, Value %d\n", key, value)
	}

	// another way to create a map
	m2 := map[string]int{
		"x": 1,
		"y": 2,
	}

	delete(m2, "y")

	for k, v := range m2 {
		fmt.Printf("Key %s, Value %d\n", k, v)
	}

}
