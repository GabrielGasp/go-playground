package main

import (
	"fmt"
)

func main() {
	x(nil)
}

func x(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
