package main

import "fmt"

func message(message string) {
	fmt.Println(message)
}

func randomValues(a int, b float64, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	message("Hi I'm a message")
	randomValues(5, 5.05, "Hello World")

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(3)
	fmt.Println(value1, value2)
}
