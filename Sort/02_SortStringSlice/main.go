package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Durgesh", "Pooja", "Armaan"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
