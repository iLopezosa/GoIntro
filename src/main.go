package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(txt string, c chan<- string) {
	init := time.Now()
	time.Sleep(time.Second * time.Duration(rand.Float64()*5))
	txt = fmt.Sprintf("%s %s", txt, time.Since(init))
	c <- txt
}

func main() {
	c := make(chan string, 3)

	fmt.Println("Hi")

	go say("bye1", c)
	go say("bye2", c)
	go say("bye3", c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
