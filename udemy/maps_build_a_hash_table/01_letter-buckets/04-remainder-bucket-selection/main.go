package main

import "fmt"

func main() {
	for i := 64; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
