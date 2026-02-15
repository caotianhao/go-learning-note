package main

import (
	"fmt"
	"strconv"
)

func dayOfYear(date string) int {
	yearTemp, _ := strconv.ParseInt(date[:4], 10, 64)
	monthTemp, _ := strconv.ParseInt(date[5:7], 10, 64)
	dayTemp, _ := strconv.ParseInt(date[8:], 10, 64)
	year, month, day := int(yearTemp), int(monthTemp), int(dayTemp)
	if month == 1 {
		return day
	} else if month == 2 {
		return 31 + day
	} else if month == 3 {
		if isRun(year) {
			return 60 + day
		} else {
			return 59 + day
		}
	} else if month == 4 {
		if isRun(year) {
			return 91 + day
		} else {
			return 90 + day
		}
	} else if month == 5 {
		if isRun(year) {
			return 121 + day
		} else {
			return 120 + day
		}
	} else if month == 6 {
		if isRun(year) {
			return 152 + day
		} else {
			return 151 + day
		}
	} else if month == 7 {
		if isRun(year) {
			return 182 + day
		} else {
			return 181 + day
		}
	} else if month == 8 {
		if isRun(year) {
			return 213 + day
		} else {
			return 212 + day
		}
	} else if month == 9 {
		if isRun(year) {
			return 244 + day
		} else {
			return 243 + day
		}
	} else if month == 10 {
		if isRun(year) {
			return 274 + day
		} else {
			return 273 + day
		}
	} else if month == 11 {
		if isRun(year) {
			return 305 + day
		} else {
			return 304 + day
		}
	} else {
		if isRun(year) {
			return 335 + day
		} else {
			return 334 + day
		}
	}
}

func isRun(n int) bool {
	return (n%4 == 0 && n%100 != 0) || n%400 == 0
}

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
	//fmt.Println(isRun(2000))
	//fmt.Println(isRun(1900))
	//fmt.Println(isRun(1904))
}
