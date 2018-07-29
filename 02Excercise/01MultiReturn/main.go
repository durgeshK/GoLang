package main

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	quot, isEven := half(1)
	println(quot, isEven)
	quot, isEven = half(2)
	println(quot, isEven)
	quot, isEven = half(5)
	println(quot, isEven)
	quot, isEven = half(21)
	println(quot, isEven)
}
