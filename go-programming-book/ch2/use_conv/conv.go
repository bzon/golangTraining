package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bzon/golangTraining/go-programming-book/ch2/conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, _ := strconv.ParseFloat(arg, 64)
		fmt.Println(conv.Meter(num))
	}

}
