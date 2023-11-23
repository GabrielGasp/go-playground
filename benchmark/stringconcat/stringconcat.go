package stringconcat

import "fmt"

func WithSprintf() string {
	s := fmt.Sprintf("%s-%s", "Hello", "World")

	return s
}

func WithPlus() string {
	s := "Hello-" + "World"

	return s
}
