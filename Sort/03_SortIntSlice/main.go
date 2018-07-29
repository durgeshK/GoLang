package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{23, 25, 17, 14, 10, 55, 108}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	//Reverse sort
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	println(n)
}
