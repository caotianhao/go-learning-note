package main

import "fmt"

func numberOfDays(year, month int) int {
	isRun := func(y int) bool {
		if y%100 != 0 {
			return y%4 == 0
		} else {
			return y%400 == 0
		}
	}
	if month == 1 || month == 3 || month == 5 || month == 7 ||
		month == 8 || month == 10 || month == 12 {
		return 31
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else if month == 2 && isRun(year) {
		return 29
	}
	return 28
}

func main() {
	fmt.Println(numberOfDays(1999, 12))
}
