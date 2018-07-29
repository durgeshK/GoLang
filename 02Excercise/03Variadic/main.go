package main

func findGreatest(args ...int) int {
	maximum := -99999
	for _, num := range args {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func main() {
	maxValue := findGreatest(1, 5, 90, 0, 12, 43, 9, 25)
	//maxValue := findGreatest()
	println(maxValue)
}
