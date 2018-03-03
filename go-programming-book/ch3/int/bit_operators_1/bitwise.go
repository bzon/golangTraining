package main

import "fmt"

func main() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Println("------------------------------------------------")
	fmt.Printf("x   | %08b | %d\n", x, x)
	fmt.Printf("y   | %08b | %d\n", y, y)
	fmt.Println("------------------------------------------------")

	fmt.Printf("x: %08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("y: %08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("x&y: %08b\n", x&y)   // "00000010", the intersection {1}
	fmt.Printf("x|y: %08b\n", x|y)   // "00100110", the union {1, 2, 5}
	fmt.Printf("x^y: %08b\n", x^y)   // "00100100", the symmetric difference {2, 5}
	fmt.Printf("x&^y: %08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Printf("%08b&(1<<%08b): %d\n", x, i, i) // "1", "5"
		}
	}

	fmt.Printf("x<<1: %08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("x>>1: %08b\n", x>>1) // "00010001", the set {0, 4}

}
