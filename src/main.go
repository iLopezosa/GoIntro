package main

import (
	"fmt"
)

type car struct {
	brand string
	year  uint16
}

func main() {
	myCar := car{brand: "Toyota", year: 1995}
	fmt.Println(myCar)

	// Alt
	var otherCar car
	otherCar.brand = "Honda"
	fmt.Println(otherCar)
}
