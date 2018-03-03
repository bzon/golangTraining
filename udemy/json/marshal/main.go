package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

type Person2 struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"james", "bond", 20, 007}
	bs, _ := json.Marshal(p1) // returns a slice of bytes
	fmt.Println(string(bs))

	p2 := Person2{"james", "bond", 20}
	bs1, _ := json.Marshal(p2)
	fmt.Println(string(bs1))

}
