package main

import (
	"fmt"
)

type teste int

func main() {
	x := 5
	var y teste = teste(x)

	fmt.Println(x, y)
}
