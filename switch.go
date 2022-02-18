package main

import "fmt"

func main() {

	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Is pair")
	default:
		fmt.Println("Is odd")
	}
}
