package main

import "fmt"

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	fmt.Println(array)
	fmt.Println(len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	// Slice options
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:5])
	fmt.Println(slice[4:])

	// Add Slice element
	slice = append(slice, 7)
	fmt.Println(slice)

	// Add slice new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
