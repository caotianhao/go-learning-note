package main

import (
	"fmt"
	"strconv"
)

func digitCount(num string) bool {
	l, cntMap, flag := len(num), make(map[int]int), true
	for i := 0; i < l; i++ {
		temp, _ := strconv.ParseInt(string(num[i]), 10, 64)
		cntMap[int(temp)]++
	}
	fmt.Println(cntMap)
	for i := 0; i < l; i++ {
		if int(num[i]-48) != cntMap[i] {
			flag = false
		}
	}
	return flag
}

func main() {
	fmt.Println(digitCount("1210"))
}
