package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func say(txt string, wg *sync.WaitGroup) {
	init := time.Now()

	defer wg.Done()

	time.Sleep(time.Second * time.Duration(rand.Float64()*5))
	fmt.Println(txt, time.Since(init))
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hi")
	wg.Add(3)

	go say("World 1", &wg)
	go say("World 2", &wg)
	go say("World 3", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("bye")
	go func(text string) {
		fmt.Println(text)
	}("bye")
	go func(text string) {
		fmt.Println(text)
	}("bye")

	time.Sleep(time.Second * 1)
}
