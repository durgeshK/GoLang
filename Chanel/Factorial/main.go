package main

import "fmt"

func main() {
	ch := fact(4)
	res := read(ch)
	fmt.Println(<-res)
}

func read(c chan int) chan int {
	ch := make(chan int)

	go func() {
		product := 1

		for n := range c {
			product *= n
		}

		ch <- product
	}()

	return ch
}

func fact(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
