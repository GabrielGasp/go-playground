package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) setAge(age int) {
	p.Age = age
}

func main() {
	person1 := &Person{
		Name: "John",
		Age:  30,
	}

	fmt.Println(person1.Age) // 33
	person1.setAge(40)
	fmt.Println(person1.Age) // 40
}
