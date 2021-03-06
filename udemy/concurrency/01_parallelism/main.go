package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func print(word string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(word, i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go print("foo")
	go print("bar")
	wg.Wait()
}
