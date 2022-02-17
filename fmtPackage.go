package main

import "fmt"

func main() {
	// Println
	helloMessage := "hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Platzi"
	courses := 500

	fmt.Printf("%s have more than %d courses\n", name, courses)
	fmt.Printf("%v have more than %v courses\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%s have more than %d courses", name, courses)
	fmt.Println(message)

	// Data type with FMT
	fmt.Printf("helloMessage: %T\n", helloMessage)
}
