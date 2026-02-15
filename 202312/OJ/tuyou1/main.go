package main

import "fmt"

func longestSameChar(s string) int {
	if len(s) == 1 {
		return 1
	}
	hm := make(map[rune]struct{})
	for _, v := range s {
		hm[v] = struct{}{}
	}
	if len(hm) == 1 {
		return len(s)
	}
	res := -1
	ss := s + s
	l := len(ss)
	left, right := 0, 1
	for right < l {
		if ss[right] == ss[left] {
			right++
		} else {
			res = max(res, right-left)
			left = right
			right++
		}
	}
	return res
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	fmt.Println(longestSameChar(s))
}
