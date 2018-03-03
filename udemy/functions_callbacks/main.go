package main

import (
	"fmt"
)

func greetings(name string, callback func(string)) {
	callback(name)
}

func main() {
	greetings("Bryan", func(s string) {
		fmt.Println("Greetings!", s)
	})
	filtered := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n != 5 // change this for testing!
	})
	fmt.Println(filtered)
}

func filter(numbers []int, callback func(int) bool) []int {
	filtered := []int{}
	for _, v := range numbers {
		if callback(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// callback: passing a func as an argument
// not really idiomatic ..
