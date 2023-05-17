package main

import (
	"fmt"
)

func message(txt string, c chan string) {
	c <- txt
}

func main() {
	c := make(chan string, 2)
	c <- "Msg 1"
	c <- "Msg 2"

	fmt.Println(len(c), cap(c))

	// Range & close
	close(c)

	// c <- "Msg 3"

	for msg := range c {
		fmt.Println(msg)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("msg1", email1)
	go message("msg2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("From email1:", m1)
		case m2 := <-email2:
			fmt.Println("From email2:", m2)
		}
	}

	// fmt.Println("Hi")

	// go say("bye1", c)
	// go say("bye2", c)
	// go say("bye3", c)

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}
