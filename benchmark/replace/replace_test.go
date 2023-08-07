package replace_test

import (
	"bench/replace"
	"fmt"
	"testing"
)

func BenchmarkReplaceWithRegex(b *testing.B) {
	mapSize := 300

	str := "User {{_p_name}} is {{_p_age}} years old"

	params := map[string]string{
		"name": "John",
		"age":  "30",
	}

	for i := 0; i < mapSize; i++ {
		params[fmt.Sprintf("%d", i)] = fmt.Sprintf("Value %d", i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		replace.ReplaceWithRegex(str, params)
	}
}

func BenchmarkReplaceWithFor(b *testing.B) {
	mapSize := 300

	str := "User {{_p_name}} is {{_p_age}} years old"

	params := map[string]string{
		"name": "John",
		"age":  "30",
	}

	for i := 0; i < mapSize; i++ {
		params[fmt.Sprintf("%d", i)] = fmt.Sprintf("Value %d", i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		replace.ReplaceWithFor(str, params)
	}
}
