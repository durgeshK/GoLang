package main

import (
	"fmt"
	"math/rand"
)

func fact(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			out <- rand.Intn(25)
		}
		close(out)
	}()
	return out
}

func calculateFact(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- fmt.Sprintf("Factorial %d : %d", n, fact(n))
		}
		close(out)
	}()
	return out
}

func main() {
	out := calculateFact(generate())
	for n := range out {
		fmt.Println(n)
	}
}
