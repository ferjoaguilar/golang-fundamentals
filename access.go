package main

import (
	"fmt"

	mypackage "golang/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	mypackage.PrintMessage()
}
