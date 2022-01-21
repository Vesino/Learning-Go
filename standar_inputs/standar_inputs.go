package main

import "fmt"

func main() {
	var input string
	var number int

	fmt.Println("Provide an input")
	fmt.Scanln(&input)
	fmt.Println("Provide a number")
	fmt.Scanln(&number)

	fmt.Printf("Your standar input was: %s, %d \n", input, number)
}
