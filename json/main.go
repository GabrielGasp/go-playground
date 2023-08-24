package main

import (
	"encoding/json"
	"fmt"
)

type jstruct struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

func main() {
	j := []byte(`{"name":"Xablau","age":null}`)

	s := jstruct{}

	json.Unmarshal(j, &s)

	fmt.Printf("%+v\n", s)

	jj := jstruct{
		Name: "Xena",
	}

	jj2, _ := json.Marshal(jj)

	fmt.Println(string(jj2))
}
