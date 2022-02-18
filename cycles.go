package main

import "fmt"

func main() {
	// Conditional 'for'
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For While
	counter := 0

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
