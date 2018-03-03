package main

import "fmt"

func main() {
	// Print hexadecimal
	fmt.Printf("%#x\n", "hello")
	fmt.Println(string(99))
	fmt.Printf("%T %d\n", 'a', 'a')
	fmt.Printf("%T %s\n", "a", "a")
	fmt.Println(`
	hey ho
	hay he`)
	// emojis
	x := '💪'                                     // int32
	fmt.Printf("%s  type is %T\n", string(x), x) //
	// hello in bytes
	bs := []byte{104, 101, 108, 108, 111}
	fmt.Println(string(bs))

	fmt.Printf("%T\n", "hello"[1])
}
