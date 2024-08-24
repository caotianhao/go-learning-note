package main

import (
	"fmt"
	"strings"
)

func candy(s string, k int) (res int) {
	hm := make(map[int]struct{})
	hm[0] = struct{}{}
	cur := 0
	s = strings.Repeat(s, k)
	for _, v := range s {
		if v == 'R' {
			cur++
			if _, ok := hm[cur]; !ok {
				hm[cur] = struct{}{}
				res += abs(cur)
			}
		} else {
			cur--
			if _, ok := hm[cur]; !ok {
				hm[cur] = struct{}{}
				res += abs(cur)
			}
		}
	}
	return
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(candy("RRR", 2))
}
