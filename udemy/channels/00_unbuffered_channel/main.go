package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered
	c := make(chan int)
	// buffered
	// c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
