package main

import "fmt"

func main() {
	var stringArray [58]string
	fmt.Println(stringArray)

	for x := 65; x < 122; x++ {
		stringArray[x-65] = string(x)
	}
	fmt.Println(stringArray)
}
