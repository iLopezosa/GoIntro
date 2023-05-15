package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	val1 := 1
	val2 := 2

	if val1 == 1 {
		fmt.Println("It is 1")
	} else {
		fmt.Println("It is NOT 1")
	}

	// AND
	if val1 == 1 && val2 == 2 {
		fmt.Println("It is TRUE, AND")
	}

	if val1 == 0 || val2 == 3 {
		fmt.Println("It is TRUE, OR")
	}

	// Convert text to integer
	value, err := strconv.Atoi("69t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
