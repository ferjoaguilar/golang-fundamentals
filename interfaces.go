package main

import "fmt"

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base float64
	tall float64
}

func (c square) area() float64 {
	return c.base * c.base
}

func (r rectangle) area() float64 {
	return r.base * r.tall
}

func calculate(f figures2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, tall: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// List interfaces

	myInterface := []interface{}{"Hola", 12, 4.90}

	fmt.Println(myInterface)
}
