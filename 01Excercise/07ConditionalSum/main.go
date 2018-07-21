package main

func main() {
	startNumber := 1
	endNumber := 1000000
	sum := 0
	for startNumber < endNumber {
		if startNumber%3 == 0 || startNumber%5 == 0 {
			sum = sum + startNumber
		}
		startNumber++
	}
	println("Sum is", sum)
}
