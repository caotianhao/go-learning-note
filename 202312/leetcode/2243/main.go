package main

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		str := ""
		for i := 0; i < len(s)/k; i++ {
			str += help2243(s[k*i : k*i+k])
		}
		if len(s)%k != 0 {
			str += help2243(s[len(s)/k*k:])
		}
		s = str
	}
	return s
}

func help2243(str string) string {
	sum := 0
	for _, v := range str {
		t, _ := strconv.ParseInt(string(v), 10, 64)
		sum += int(t)
	}
	return strconv.FormatInt(int64(sum), 10)
}

func main() {
	fmt.Println(digitSum("11111222223", 3))
	fmt.Println(digitSum("5456432165431654", 2))
	//                      9117311779  1084816  11296  2116  37
}
