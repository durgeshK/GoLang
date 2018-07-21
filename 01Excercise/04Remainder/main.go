package main

import "fmt"

func main() {
	var smallInteger int
	var largeInteger int
	fmt.Print("Please enter the larger integer: ")
	fmt.Scan(&largeInteger)
	fmt.Print("Please enter the smaller integer: ")
	fmt.Scan(&smallInteger)
	if smallInteger < largeInteger {
		fmt.Printf("\nRemainder is :%d\n", largeInteger%smallInteger)
	} else {
		fmt.Println("Bad input")
	}
}
