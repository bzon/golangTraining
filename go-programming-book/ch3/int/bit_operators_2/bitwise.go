// Bits computation
// Good reference: https://www.youtube.com/watch?v=d0AwjSpNXR0

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	operators := []string{"&", "|", "^", "&^", "<<", ">>"}
	for _, operator := range operators {
		bitwiseOperator(uint8(rand.Intn(255)), uint8(rand.Intn(255)), operator)
	}
}

func printBits(a, b uint8) {
	fmt.Printf("a: %d in binary is %08b\n", a, a)
	fmt.Printf("b: %d in binary is %08b\n", b, b)
}

func bitwiseOperator(a, b uint8, operator string) {
	fmt.Println("---------------------------------")
	printBits(a, b)
	var x uint8
	switch operator {
	case "&":
		fmt.Println("Using bitwise AND https://youtu.be/d0AwjSpNXR0?t=129")
		x = a & b
	case "|":
		fmt.Println("Using bitwise OR https://youtu.be/d0AwjSpNXR0?t=272")
		x = a & b
	case "^":
		fmt.Println("Using bitwise Exclusive OR (XOR) https://youtu.be/d0AwjSpNXR0?t=340")
		x = a ^ b
	case "&^":
		fmt.Println("Using Golang bitclear AND NOT https://stackoverflow.com/a/34459527")
		x = a &^ b
	case "<<":
		fmt.Println("Using bitwise LEFT SHIFT https://youtu.be/d0AwjSpNXR0?t=419")
		x = a << b
	case ">>":
		fmt.Println("Using bitwise RIGHT SHIFT https://youtu.be/d0AwjSpNXR0?t=482")
		x = a >> b
	default:
		os.Exit(1)
	}
	fmt.Printf("bitwise computation, <a %s b> in binary is %08b or %d in decimal\n", operator, x, x)
	fmt.Println("---------------------------------")
	fmt.Printf("a: %08b // %d\n", a, a)
	fmt.Printf("b: %08b // %d\n", b, b)
	fmt.Printf("x: %08b // %d\n", x, x)
}
