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

func (s *stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *stack[T]) Pop() (T, error) {
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

	s.Push(1)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)
	s.Push(3)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
}
