package main

import "fmt"

func main() {
	const hi = "Hello"
	const world = "World!"

	// Println()
	fmt.Println(hi, world)
	fmt.Println(hi, world)

	// Printf()
	platzi := "Platzi"
	courses := 500
	fmt.Printf("%s has more than %d courses\n", platzi, courses)
	fmt.Printf("%v has more than %d courses\n", platzi, courses)

	// Sprintf()
	message := fmt.Sprintf("%s has more than %d courses", platzi, courses)
	fmt.Println(message)

	// Data types
	fmt.Printf("message: %T\n", message)
	fmt.Printf("courses: %T\n", courses)

}
