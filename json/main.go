package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type jstruct struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

func main() {
	j := []byte(`{"name":27,"age":null}`)

	s := jstruct{}

	err := json.Unmarshal(j, &s)

	if err != nil {
		var typeError *json.UnmarshalTypeError
		if errors.As(err, &typeError) {
			fmt.Printf("Invalid type for field %q, expected a %q but got a %q.\n", typeError.Field, typeError.Type, typeError.Value)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Printf("%+v\n", s)

	// jj := jstruct{
	// 	Name: "Xena",
	// }

	// jj2, _ := json.Marshal(jj)

	// fmt.Println(string(jj2))
}
