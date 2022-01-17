package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 9
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("Palomita", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Maps are data type of the kind: key & value:
	m := make(map[string]int)
	m["Key1"] = 6
	fmt.Println(m["Key1"])

	// Slices are data type similar to an array
	s := []int{1, 2, 3}
	for _, value := range s {
		fmt.Println(value)
	}
	s = append(s, 13)
	for _, value := range s {
		fmt.Println(value)
	}

	// Punteros
	// g := 27
	// fmt.Println(g)
	// h := &g
	// fmt.Println(h)
	// fmt.Println(*h)

	// This subroutin will execute in another routine due the reserved word go
	// But in order that the main routin execute the following a channel must be created and relate it to our subroutin
	c := make(chan int)
	go doSomething(c)
	go doOtherThing()
	go doSomething(c)
	<-c
	fmt.Println("Aqui llegare hasta que las rutinas en los channels terminen? No necesariamente....")
	fmt.Println("Lo imprimirá?")

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	println("Do something outside of principal :v")
	c <- 1
}

func doOtherThing() {
	println("Me imprimo primero, me ejecuto en el principal")
}
