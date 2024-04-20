package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	startDate := time.Date(2023, 8, 23, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC)

	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		fmt.Println(date.Format("2006-01-02"))
	}

	fmt.Println("Time taken:", time.Since(now))
}
