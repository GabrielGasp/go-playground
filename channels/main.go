package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 1; i < 1000; i++ {
			c <- i
			fmt.Printf("Add %d to channel\n", i)
		}

		close(c)
	}()

	for n := range c {
		fmt.Printf("Read %d from channel\n", n)
	}
}
