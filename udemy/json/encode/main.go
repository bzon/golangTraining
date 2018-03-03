package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	json.NewEncoder(os.Stdout).Encode(p1)
}
