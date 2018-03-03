package main

import (
	"bufio"
	"fmt"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert in to ByteCounter
	return len(p), nil
}

func main() {
	fmt.Printf("Words count: %d\n", countWords())
	fmt.Printf("Lines count: %d\n", countLines())
}

func countLines() int {
	file, _ := os.Open("/tmp/lines.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lc := 0
	for scanner.Scan() {
		lc++
	}
	return lc
}

func countWords() int {
	file, _ := os.Open("/tmp/lines.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
