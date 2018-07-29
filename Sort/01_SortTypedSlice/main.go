package main

import (
	"fmt"
	"sort"
)

type person []string

func main() {
	s := person{"Durgesh", "Pooja", "Armaan", "Kumar", "Kashyap", "Kumari"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
