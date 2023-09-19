package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "UTC")

	startAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	finishAt := time.Date(2023, 1, 7, 23, 59, 59, 999999, time.UTC)

	startDate := startAt.Truncate(24 * time.Hour)
	finishDate := finishAt.Truncate(24 * time.Hour).Add(24 * time.Hour)

	fmt.Println(startDate)
	fmt.Println(finishDate)
	fmt.Println(finishDate.Sub(startDate))
}
