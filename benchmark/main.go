package main

import (
	"bench/replace"
	"fmt"
)

func main() {
	str := "User {{_p_name}} is {{_p_age}} years old"

	params := map[string]string{
		"name": "John",
		"age":  "30",
	}

	// Replace with regex
	rgx := replace.ReplaceWithRegex(str, params)
	fmt.Println(rgx)

	// Replace with for
	f := replace.ReplaceWithFor(str, params)
	fmt.Println(f)
}
