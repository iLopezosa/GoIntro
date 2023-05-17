package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("PC: RAM %d GB, Disk space %d GB, Brand %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 8, disk: 512, brand: "asus"}
	fmt.Println(myPC)

}
