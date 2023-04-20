package main

import "fmt"

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

type embedded struct {
	Dog
}

func main() {
	e := embedded{}
	e.Speak()
}
