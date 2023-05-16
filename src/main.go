package main

import (
	"pk/mypackage"

	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Honda"
	fmt.Println(myCar)
}
