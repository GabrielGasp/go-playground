package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayMyName() {
	fmt.Println(p.Name)
}

type NotAPerson struct {
	Name string
}

func main() {
	var person any = Person{"John"}

	// This will work \/
	assertedPerson1 := person.(Person)
	assertedPerson1.SayMyName()

	// This will panic \/
	assertedPerson2 := person.(NotAPerson)
	fmt.Println(assertedPerson2.Name)
}
