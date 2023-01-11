package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name *string `json:"name,omitempty"`
	Age  *int    `json:"age,omitempty"`
}

func main() {
	JSON := `{"name": "John", "age": 0}`
	p := Person{}

	// Unmarshal JSON to Person
	if err := json.Unmarshal([]byte(JSON), &p); err != nil {
		panic(err)
	}

	// print values
	fmt.Println("Struct:", p)
	fmt.Println("Name:", *p.Name)
	fmt.Println("Age:", *p.Age)

	// Marshal Person to JSON
	if b, err := json.Marshal(p); err != nil {
		panic(err)
	} else {
		fmt.Println("JSON:", string(b))
	}
}
