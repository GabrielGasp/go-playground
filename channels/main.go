package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		for i := 1; i <= 1000; i++ {
			c <- i
			fmt.Printf("Add %d to channel\n", i)
			time.Sleep(time.Millisecond * 1)
		}
		close(c)
		wg.Done()
	}()

	go func() {
		for n := range c {
			fmt.Printf("Read %d from channel\n", n)
		}
		wg.Done()
	}()

	wg.Wait()
}
