package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func sumOf[T Number](values []T) T {
	var sum T

	for _, v := range values {
		sum += v
	}

	return sum
}

func main() {
	values := []int64{1, 2, 3, 4, 5}
	fmt.Println(sumOf(values))

	values2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumOf(values2))
}
