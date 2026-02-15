package main

import (
	"fmt"
	"strconv"
)

func oddString(words []string) string {
	l, mm := len(words[0]), map[string][]string{}
	for _, v := range words {
		str := ""
		for i := 0; i < l-1; i++ {
			str += strconv.FormatInt(int64(int(v[i+1])-int(v[i])), 10)
			str += "a"
		}
		mm[str] = append(mm[str], v)
	}
	for _, v := range mm {
		if len(v) == 1 {
			return v[0]
		}
	}
	return ""
}

func main() {
	fmt.Println(oddString([]string{"adc", "wzy", "abc"}))
	fmt.Println(oddString([]string{"aaa", "bob", "ccc", "ddd"}))
	fmt.Println(oddString([]string{"az", "za", "az"}))
	fmt.Println(oddString([]string{"abm", "bcn", "alm"}))
}
