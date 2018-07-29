package main

import "fmt"

func functionDefination() func(int) (int, bool) {
	return func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
}

func main() {
	function := functionDefination()
	fmt.Printf("%T", function)
	quot, isEven := function(1)
	println(quot, isEven)
	quot, isEven = function(2)
	println(quot, isEven)
	quot, isEven = function(5)
	println(quot, isEven)
	quot, isEven = function(21)
	println(quot, isEven)
}
