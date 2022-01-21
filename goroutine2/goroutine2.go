package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Wait in order to not exit the main function before the go routines finish
var wg sync.WaitGroup

func f() {
	for i := 0; i < 10; i++ {
		fmt.Println(":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	wg.Done()

}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
}
