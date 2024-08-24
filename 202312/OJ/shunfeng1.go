package main

import (
	"fmt"
	"strconv"
)

func calNM64(n, m int64) (r string) {
	ans := n * m
	s := strconv.FormatInt(ans, 10)
	l := len(s)
	if l%3 == 1 {
		s = "00" + s
	} else if l%3 == 2 {
		s = "0" + s
	}
	for i := 0; i < len(s); i++ {
		if i%3 == 0 {
			r += " "
		}
		r += string(s[i])
	}
	return r
}

func main() {
	var n, m int64
	_, _ = fmt.Scanf("%d %d", &n, &m)
	fmt.Printf("%s", calNM64(n, m))
}
