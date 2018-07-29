package main

func main() {
	//Should be true
	value := (true && false) || (false && true) || !(false && false)
	println(value)
}
