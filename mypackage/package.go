package mypackage

import "fmt"

// Public access
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate
type carPrivate struct {
	brand string
	year  int
}

// Public function
func PrintMessage() {
	fmt.Println("Hiiiii!!!!")
}
