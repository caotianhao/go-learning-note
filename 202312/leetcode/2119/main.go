package main

import (
	"fmt"
	"strconv"
)

func isSameAfterReversals(num int) bool {
	rev0 := num
	rev1, _ := strconv.ParseInt(reverse2119(strconv.FormatInt(int64(num), 10)), 10, 64)
	rev2, _ := strconv.ParseInt(reverse2119(strconv.FormatInt(rev1, 10)), 10, 64)
	if rev0 == int(rev2) {
		return true
	} else {
		return false
	}
}

func reverse2119(s string) string {
	str := []rune(s)
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str)
}

func main() {
	fmt.Println(isSameAfterReversals(526))
}
