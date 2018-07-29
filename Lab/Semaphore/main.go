package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 5
	c := make(chan int)
	done := make(chan bool)

	for j := 0; j < n; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Printf("%d ", n)
	}

}
