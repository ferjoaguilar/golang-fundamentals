package main

import "fmt"

func main() {
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("Is 1")
	} else {
		fmt.Println("Is not 1")
	}

	// With And
	if value1 == 1 && value2 == 2 {
		fmt.Println("It's true")
	}

	// With Or
	if value1 == 1 || value2 == 0 {
		fmt.Println("It's true")
	}

}
