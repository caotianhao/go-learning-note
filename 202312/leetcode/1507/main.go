package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	tmp := strings.Split(date, " ")
	tmpMap := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	if tmp[0][1] > 'a' {
		tmp[0] = "0" + tmp[0][:1]
	} else {
		tmp[0] = tmp[0][:2]
	}
	return tmp[2] + "-" + tmpMap[tmp[1]] + "-" + tmp[0]
}

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
}
