package main

import "fmt"

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 8)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	fmt.Println(slice1[0])
}
