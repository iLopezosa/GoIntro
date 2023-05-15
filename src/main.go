package main

import "fmt"

func hi(msg string) {
	fmt.Println(msg)
}

func threeArgs(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnVal(a int) int {
	return a * 2
}

func return2Vals(a int) (c, d int) {
	return a, a * 2
}

func main() {
	hi("Hello World")
	threeArgs(1, 2, "hi")
	fmt.Println("Value:", returnVal(4))
	val1, val2 := return2Vals(3)
	fmt.Println("Values:", val1, val2)
}
