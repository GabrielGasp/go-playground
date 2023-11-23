package main

import (
	"encoding/json"
	"fmt"
)

type teste struct {
	Name   string      `json:"name"`
	Number json.Number `json:"number"`
}

func main() {
	j := `{"name":"test","number":1234567890123456789012345678901234567890}`

	var t teste
	err := json.Unmarshal([]byte(j), &t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Number.String())
}
