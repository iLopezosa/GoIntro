package main

import (
	"fmt"
	"time"
)

func longFunction() {
	time.Sleep(4 * time.Second)
	fmt.Println("Hi")
}

func main() {
	// Defer
	defer longFunction()
	time.Sleep(2 * time.Second)
	fmt.Println("world")

	// Continue and break
	for i := 0; i < 10; i++ {
		// Continue
		if i == 2 {
			fmt.Println("It's 2")
			continue
		}

		if i == 8 {
			fmt.Println("BREAK")
			break
		}

		fmt.Println(i)
	}
}
