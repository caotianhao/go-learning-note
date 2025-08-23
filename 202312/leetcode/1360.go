package main

import (
	"fmt"
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	if date1 == date2 {
		return 0
	}
	a, _ := time.Parse("2006-01-02", date1)
	b, _ := time.Parse("2006-01-02", date2)
	aa := time.Since(a).Hours() / 24
	bb := time.Since(b).Hours() / 24
	k := math.Abs(aa - bb)
	if math.Abs(math.Ceil(k)-k) < math.Abs(math.Floor(k)-k) {
		return int(math.Ceil(k))
	}
	return int(math.Floor(k))
}

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2024-11-07", "2057-03-09"))
}
