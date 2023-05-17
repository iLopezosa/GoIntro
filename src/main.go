package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplidateRAM() {
	myPC.ram *= 2
}

func main() {
	a := 50
	b := &a
	c := b

	fmt.Println(a, &a, b, *b, *c)

	*b = 100
	fmt.Println(a, &a, b, *b, *c)

	myPC := pc{ram: 8, disk: 512, brand: "asus"}
	fmt.Println(myPC)

	myPC.ping()
	myPC.duplidateRAM()
	fmt.Println(myPC)
}
