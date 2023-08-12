package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Date(2023, 8, 12, 15, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 8, 12, 15, 0, 0, 0, time.Local)

	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time1.Before(time2))
	fmt.Println(time2.Sub(time1))
}
