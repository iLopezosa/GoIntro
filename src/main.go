package main

import (
	"fmt"
)

func main() {
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	// Conditionless
	value := -200
	switch {
	case value > 100:
		fmt.Println("Greater than 100")
	case value < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("No condition")
	}
}
