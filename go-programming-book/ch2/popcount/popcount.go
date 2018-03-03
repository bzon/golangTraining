package popcount

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[%d] is %b\n", i, pc[i])
	}
}

// PopCount returns the population count (number of bits) of x.
func PopCount(x uint64) int {
	var popcount int
	for i := 0; i < 8; i++ {
		fmt.Printf("%b\n", pc[i])
		popcount += int(pc[byte(x>>(uint64(i)*8))])
	}
	return popcount
	// return int(pc[byte(x>>(0*8))] +
	// 	pc[byte(x>>(1*8))] +
	// 	pc[byte(x>>(2*8))] +
	// 	pc[byte(x>>(3*8))] +
	// 	pc[byte(x>>(4*8))] +
	// 	pc[byte(x>>(5*8))] +
	// 	pc[byte(x>>(6*8))] +
	// 	pc[byte(x>>(7*8))])
}
