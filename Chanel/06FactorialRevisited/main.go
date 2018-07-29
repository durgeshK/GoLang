package main

import (
	"fmt"
	"math/rand"
	"sync"
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
		for i := 1; i <= 100000; i++ {
			out <- rand.Intn(22)
		}
		close(out)
	}()
	return out
}

func calculateFact(funcIndex int, in <-chan int) chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- fmt.Sprintf("%d - Factorial %d : %d", funcIndex, n, fact(n))
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
	inStream := generate()
	fanOutCount := 10
	var channelList []chan string
	for i := 0; i < fanOutCount; i++ {
		channelList = append(channelList, calculateFact(i+1, inStream))
	}

	out := mergeChannels(channelList...)

	index := 0
	for n := range out {
		index++
		fmt.Println(index, n)
	}
}
