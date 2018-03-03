// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		fmt.Println("--------------------------------------")
		fmt.Println("Printing the contents of", filename)
		fmt.Println("--------------------------------------")
		fmt.Println(string(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	fmt.Println("--------------------------------------")
	fmt.Println("| Count | Word")
	fmt.Println("--------------------------------------")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("   %d\t  %s\n", n, line)
		}
	}
}

// Donovan, Alan A. A.. The Go Programming Language (Addison-Wesley Professional Computing Series) (p. 10). Pearson Education. Kindle Edition.
