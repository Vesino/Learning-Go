package main

import "fmt"

func multiplyByTwo(n int) int {
	fmt.Println(n * 2)
	return n * 2
}

// In concurrency some process could end before that the ones that starts!

func main() {
	for i := 0; i < 10; i++ {
		go multiplyByTwo(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Your input was %v\n", input)
	sameInput := input
	fmt.Printf("This is only a copy of the input %v\n", sameInput)
	var some int
	for i := 0; i < 2; i++ {
		some = i
	}
	fmt.Printf("some is %d\n", some)
	other := &input
	fmt.Println(*other)
}
