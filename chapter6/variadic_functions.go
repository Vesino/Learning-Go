package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9)) // 45

	// passing a whole slice of ints
	slice1 := []int{1, 2, 3, 45, 6, 35, 1}
	fmt.Println(add(slice1...)) // 93
}

// Variadic parameters
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
