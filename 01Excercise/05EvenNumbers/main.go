package main

func main() {
	startNumber := 1
	endNumber := 100
	for startNumber < endNumber {
		if startNumber%2 == 0 {
			println(startNumber)
		}
		startNumber++
	}
}
