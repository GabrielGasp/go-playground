package main

import (
	"errors"
	"fmt"
)

type stack[T any] struct {
	items []T
}

func NewStack[T any]() stack[T] {
	return stack[T]{}
}

func (s *stack[T]) Add(v T) {
	s.items = append(s.items, v)
}

func (s *stack[T]) Remove() (T, error) {
	var v T
	i := len(s.items) - 1

	if i < 0 {
		return v, errors.New("stack is empty")
	}

	v = s.items[i]
	s.items = s.items[:i]
	return v, nil
}

func main() {
	s := NewStack[int]()

	s.Add(1)
	fmt.Println(s)
	s.Add(2)
	fmt.Println(s)
	s.Add(3)
	fmt.Println(s)

	fmt.Println(s.Remove())
	fmt.Println(s)
	fmt.Println(s.Remove())
	fmt.Println(s)
	fmt.Println(s.Remove())
	fmt.Println(s)
	fmt.Println(s.Remove())
}
