package main

import (
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	str := strconv.Itoa(n)
	l := len(str)
	if l <= 3 {
		return str
	} else if l == 4 {
		return string(str[0]) + "." + str[1:]
	} else if l == 5 {
		return str[:2] + "." + str[2:]
	} else if l == 6 {
		return str[:3] + "." + str[3:]
	} else if l == 7 {
		return string(str[0]) + "." + str[1:4] + "." + str[4:]
	} else if l == 8 {
		return str[:2] + "." + str[2:5] + "." + str[5:]
	} else if l == 9 {
		return str[:3] + "." + str[3:6] + "." + str[6:]
	} else {
		return string(str[0]) + "." + str[1:4] + "." + str[4:7] + "." + str[7:]
	}
}

func main() {
	fmt.Println(thousandSeparator(56369))
}
