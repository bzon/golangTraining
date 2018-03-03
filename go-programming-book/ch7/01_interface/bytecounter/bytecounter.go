package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert in to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "World"
	// since c is a ByteCounter that implements io.Writer, it is acceptable
	fmt.Fprintf(&c, "%s", name)
	// so, Fprintf writes the byte size of name to &c
	fmt.Println(c)
}
