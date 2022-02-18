package main

import "fmt"

func main() {
	slice := []string{"Pear", "Banana", "Orange"}

	for i, value := range slice {
		fmt.Println(i, value)
	}
}
