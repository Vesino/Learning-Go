package main

import "fmt"

type Person struct {
	age  int
	name string
}

type Employe struct {
	id int
}

type FullTimeEmploye struct {
	Person
	Employe
}

// Go Does not allow inheritance, instead it use composition over inheritence

func main() {
	fllTimeEmploye := new(FullTimeEmploye)
	fllTimeEmploye.id = 2
	fllTimeEmploye.age = 27
	fllTimeEmploye.name = "Vesino"
	fmt.Printf("The employe id is: %d, name: %s\n", fllTimeEmploye.id, fllTimeEmploye.name)
}
