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

	layout := "Monday, January 2, 2006, 15:04 -0700"
	dateTime := "Thursday, October 23, 2023, 16:12 -0300"

	// If timezone is not set, the time will be parsed as UTC.
	// If you want to parse the time as a specific timezone, you need to include the timezone or use time.ParseInLocation.
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(parsedTime)
	fmt.Println(parsedTime.UTC())
}
