package main

import (
	"fmt"
	"strconv"
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
}
