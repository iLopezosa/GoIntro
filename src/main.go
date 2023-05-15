package main

import "fmt"

func main() {
	// Constants declaration
	const pi float64 = 3.14159265359
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables declaration
	base := 12          // infers the type of variable
	var height int = 13 // explicit type declaration
	var area int        // zero value

	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	fmt.Println(squareArea(2.12))
}

// Area of a square
func squareArea(sideLen float64) float64 {
	return sideLen * sideLen
}
