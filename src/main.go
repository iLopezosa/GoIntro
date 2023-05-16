package main

import (
	"fmt"
	"time"
)

func isPalyndromeFast(text string) bool {
	var length int = len(text) - 1

	for i := range text {
		if text[i] != text[length-i] {
			return false
		}
	}

	return true
}

func isPalyndromeSlow(text string) bool {
	var reverseText string

	for i := len(text) - 1; i >= 0; i-- {
		reverseText += string(text[i])
	}

	return reverseText == text
}

func main() {
	// slice := []string{"hi", "what's", "up"}

	// for i, value := range slice {
	// 	fmt.Printf("The characters of the %d:%s string are of type %T\n", i, value, slice[i][0])
	// }
	start := time.Now()
	fmt.Println(isPalyndromeFast("amor a roma"))
	fmt.Println(isPalyndromeFast("ama"))
	fmt.Println(isPalyndromeFast("ieanorstienariostenaorisetnaiorsetnaioe"))
	fmt.Println(isPalyndromeFast("aiorsetnarioarsarstarstoaiersntoiearnsoitenarsioteiarnsoeitniarensotieanrsioetnaioresntoiearnsiotenariosetnaioresntioaernsiotenariosetnarioesntioarenstioeanrsiotenarioestnoiaresntioarenstoiaernstioearnsiotenarsioetnaorisetnioaresntioearnsoitenarsiotenaroisetnoiaresntoiaernstioeanrssetnaoiresntoaiersntoiaernstoiearnsoitenarsioetnariosetnioaresntioearnsiotenarsiotenariosetnarsktgienwkfein"))
	fmt.Println(isPalyndromeFast("arst arst arst tsra tsra tsra"))
	fmt.Printf("Fast time: %d\n", time.Since(start))
	start = time.Now()
	fmt.Println(isPalyndromeSlow("amor a roma"))
	fmt.Println(isPalyndromeSlow("ama"))
	fmt.Println(isPalyndromeSlow("ieanorstienariostenaorisetnaiorsetnaioe"))
	fmt.Println(isPalyndromeFast("aiorsetnarioarsarstarstoaiersntoiearnsoitenarsioteiarnsoeitniarensotieanrsioetnaioresntoiearnsiotenariosetnaioresntioaernsiotenariosetnarioesntioarenstioeanrsiotenarioestnoiaresntioarenstoiaernstioearnsiotenarsioetnaorisetnioaresntioearnsoitenarsiotenaroisetnoiaresntoiaernstioeanrssetnaoiresntoaiersntoiaernstoiearnsoitenarsioetnariosetnioaresntioearnsiotenarsiotenariosetnarsktgienwkfein"))
	fmt.Println(isPalyndromeSlow("arst arst arst tsra tsra tsra"))
	fmt.Printf("Slow time: %d\n", time.Since(start))
}
