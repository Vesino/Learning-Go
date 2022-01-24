package main

import (
	"fmt"
	"time"
)

func main() {
	// Anonymous functions
	z := func(n int) int {
		return n * 2
	}(5)
	fmt.Println(z)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
