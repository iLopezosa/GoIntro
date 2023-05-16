package main

import (
	"fmt"
)

func main() {
	m := make(map[string]uint8)

	m["Joseph"] = 0
	m["Francis"] = 20
	m["Anne"] = 32

	fmt.Println(m)

	// Read map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Get one value
	value, found := m["Joseph"]
	fmt.Println(value, found)
}
