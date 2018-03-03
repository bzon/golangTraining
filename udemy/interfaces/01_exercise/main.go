package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	fmt.Println("Ascending..")
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("Descending..")
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Sorting a []string using sort.StringSlice")
	// StringSlice is a type, therefore you can use it to convert
	// like Type(ToConvert) or string(65)
	sortedS := sort.StringSlice(s)
	sort.Sort(sortedS)
	fmt.Println(sortedS)
	sort.Sort(sort.Reverse(sortedS))
	fmt.Println(sortedS)

	n := []int{5, 4, 10, 8, 1}
	fmt.Println("Sorting a []int using sort.IntSlice")
	sortedN := sort.IntSlice(n)
	sort.Sort(sortedN)
	fmt.Println(sortedN)
	sort.Sort(sort.Reverse(sortedN))
	fmt.Println(sortedN)

}
