package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the moby dick text
	res, err := http.Get("http://127.0.0.1:3999/welcome/1")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create the slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBuckets(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:122])
}

func HashBuckets(word string) int {
	return int(word[0])
}
