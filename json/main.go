package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type NullableOrOptional[T any] struct {
	Present bool
	Null    bool
	Value   T
}

func (field *NullableOrOptional[T]) UnmarshalJSON(b []byte) error {
	field.Present = true

	if string(b) == "null" {
		field.Null = true
		return nil
	}

	return json.Unmarshal(b, &field.Value)
}

type jstruct struct {
	Name string                    `json:"name"`
	Age  NullableOrOptional[uint8] `json:"age"`
}

func main() {
	j := []byte(`{"name":"Xablau", "age": 10}`)

	s := jstruct{}

	err := json.Unmarshal(j, &s)

	if err != nil {
		var typeError *json.UnmarshalTypeError
		if errors.As(err, &typeError) {
			fmt.Printf("Invalid type for field %q, expected a %q but got a %q.\n", typeError.Field, typeError.Type, typeError.Value)
		} else {
			fmt.Println(err)
		}

		return
	}

	fmt.Printf("%+v\n", s)
}
