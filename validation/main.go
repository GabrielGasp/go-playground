package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func main() {
	fmt.Println(govalidator.HasUpperCase("hello"))
	fmt.Println(govalidator.HasUpperCase("Hello"))
	fmt.Println(govalidator.IsUUIDv4("hello"))
	fmt.Println(govalidator.IsUUIDv4("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))
}
