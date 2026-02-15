package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day, month, year int) string {
	return time.Now().AddDate(year-2023, month-5, day-9).Weekday().String()
}

func main() {
	// 1971 - 2100
	fmt.Println(dayOfTheWeek(12, 9, 2093))
}
