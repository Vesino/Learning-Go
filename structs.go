package main

import (
	"fmt"
)

type Employe struct {
	id       int
	name     string
	vacation bool
}

// Receiver functions as class (struct) methods
// The methods must have a pointer receiver to change the Employe value declared in the main function.
func (e *Employe) SetId(id int) {
	e.id = id
}

func (e *Employe) SetName(name string) {
	e.name = name
}

func (e *Employe) GetId() int {
	return e.id
}

func (e *Employe) GetName() string {
	return e.name
}

// Create struct (classes) with constructors:
func NewEmploye(id int, name string, vacation bool) *Employe {
	return &Employe{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// e := Employe{}
	// e.SetId(4)
	// e.name = "Name"
	// fmt.Println(e)
	// e.SetId(1)
	// fmt.Println(e)
	// e.SetName("Vesino")
	// fmt.Println(e)
	// fmt.Println(e.GetId(), e.GetName())
	e3 := NewEmploye(23, "El hobbit", true)
	fmt.Println(e3)
}
