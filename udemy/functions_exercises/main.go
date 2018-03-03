package main

import (
	"fmt"
)

func isEven(n int) (int, bool) {
	if n%2 == 0 {
		return n, true
	}
	return n, false
}

func getHighest(numbers ...float64) float64 {
	var highest float64
	for _, v := range numbers {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}

// https://projecteuler.net/problem=1
func sumOfMultiplesBelow(n int) int {
	var s int
	m := []int{}
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			m = append(m, i)
		}
	}
	fmt.Println(m)
	for _, v := range m {
		s += v
	}
	return s
}

// https://projecteuler.net/problem=2
func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b // exchanges values
		return a
	}
}

func sumFib(n int) uint64 {
	a := []int{}
	f := fibonacci()
	for i := 0; i < n; i++ {
		a = append(a, f())
	}
	// fmt.Println(a)
	var s uint64
	for _, v := range a {
		if v%2 == 0 && s < 4000000 {
			s += uint64(v)
		}
	}
	return s
}

func main() {
	// fmt.Println(isEven(4))
	// isEven2 := func(n int) (int, bool) {
	// 	if n%2 == 0 {
	// 		return n, true
	// 	}
	// 	return n, false
	// }
	// fmt.Println(isEven2(3))
	// fmt.Println(getHighest(11, 5, 121, 9, 3, 10, 8))
	// foo(1, 2)
	// foo(1, 2, 3)
	// s := []int{1, 2, 3, 4}
	// foo(s...)
	// foo()

	fmt.Println(sumFib(4000000))

	// A simpler solution ...
	// https: //siongui.github.io/2017/12/17/go-even-fibonacci-numbers-problem-2-project-euler/
	var a, b, sum uint64 = 1, 2, 2
	for a+b < 4000000 {
		next := a + b
		//fmt.Println(next)
		if next%2 == 0 {
			sum += next
		}
		a = b
		b = next
	}
	fmt.Println("the sum of the even-valued terms below 4,000,000 is", sum)
}
