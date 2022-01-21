package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	msg := <-c
	fmt.Println(msg)
	time.Sleep(time.Second * 10)
}

func main() {

}
