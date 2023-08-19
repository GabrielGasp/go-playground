package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "UTC")

	now := time.Now()

	fmt.Println(now)

	// time1 := time.Date(2023, 8, 12, 15, 0, 0, 0, time.UTC)
	// time2 := time.Date(2023, 8, 12, 12, 0, 0, 0, time.Local)

	// fmt.Println(time1)
	// fmt.Println(time2)
	// fmt.Println(time1.Before(time2))
	// fmt.Println(time2.Sub(time1))

	// t := "17-08-2023"
	// layout := "02-01-2006"
	// p, _ := time.Parse(layout, t)
	// fmt.Println(p)
}
