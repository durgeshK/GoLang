package main

func foo(args ...int) int {
	len := 0
	for _, _ = range args {
		len++
	}
	return len
}

func main() {
	println(foo(1, 2))
	println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	println(foo(aSlice...))
	println(foo())
}
