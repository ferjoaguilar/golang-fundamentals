package main

import "fmt"

func main() {
	// Const
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi", pi)
	fmt.Println("pi2", pi2)

	// Variables
	base := 12
	var tall int = 14
	var area int

	fmt.Println(base, tall, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Exercise
	// Square area
	const baseSquare = 10
	areaSquare := baseSquare * baseSquare
	fmt.Println("Square area:", areaSquare)
}
