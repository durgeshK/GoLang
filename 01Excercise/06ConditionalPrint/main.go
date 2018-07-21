package main

func main() {
	startNumber := 1
	endNumber := 100
	for startNumber < endNumber {
		if startNumber%3 == 0 && startNumber%5 == 0 {
			println("FizzBuzz")
		} else if startNumber%3 == 0 {
			println("Fizz")
		} else if startNumber%5 == 0 {
			println("Buzz")
		} else {
			println(startNumber)
		}
		startNumber++
	}
}
