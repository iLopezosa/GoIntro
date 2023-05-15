package main

import "fmt"

func main() {

	// Conditional For
	for i := 0; i <= 10; i++ {
		fmt.Print(i, ", ")
	}

	fmt.Println()

	// For while
	counter := 0
	for counter <= 10 {
		fmt.Print(counter, ", ")
		counter++
	}

	// For forever
	// foreverCounter := 0
	// for {
	// 	fmt.Print(foreverCounter, ", ")
	// 	foreverCounter++
	// }
}
