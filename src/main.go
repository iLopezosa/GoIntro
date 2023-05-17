package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	width float64
}
type rectangle struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	return s.width * s.width
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func calculate(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	mySquare := square{width: 2}
	myRectangle := rectangle{width: 2, height: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// Interface list
	myInterface := []interface{}{"Hi", 12, 3.14}
	fmt.Println(myInterface...)
}
