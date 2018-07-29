package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()

	return ch
}

func square(in chan int) chan string {
	out := make(chan string)

	go func() {
		for n := range in {
			out <- fmt.Sprintf("Square Function %d \t %d", n, n*n)
		}
		close(out)
	}()

	return out
}

func cube(in chan int) chan string {
	out := make(chan string)

	go func() {
		for n := range in {
			out <- fmt.Sprintf("Cube Function %d \t %d", n, n*n*n)
		}
		close(out)
	}()

	return out
}

func mergeChannels(in ...chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(in))

	for _, chs := range in {
		go func(ch chan string) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(chs)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	firstChannel := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	squareChannel := square(firstChannel)
	cubeChannel := cube(firstChannel)

	for str := range mergeChannels(squareChannel, cubeChannel) {
		fmt.Println(str)
	}

}
