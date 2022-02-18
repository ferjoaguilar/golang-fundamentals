package main

import "fmt"

// Struct
type game struct {
	name     string
	year     int
	category string
}

func main() {

	// Instance struct way 1
	myGame := game{
		name:     "Overwatch",
		year:     2016,
		category: "Action",
	}

	fmt.Println(myGame)

	// Instance struct way 2
	var otherGame game
	otherGame.name = "Overwatch 2"
	otherGame.category = "Action"
	fmt.Println(otherGame)
}
