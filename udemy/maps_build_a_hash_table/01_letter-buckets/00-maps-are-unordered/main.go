package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["AliceBlue"] = "#F0F8FF"
	colors["Coral"] = "#FF7F50"
	colors["DarkGray"] = "#A9A9A9"

	// Loop 3 times
	for i := 0; i < 3; i++ {
		fmt.Println(" ")
		for k, v := range colors {
			fmt.Printf("%s:%s \n", k, v)
		}
	}

}
