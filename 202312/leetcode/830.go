package main

import "fmt"

func largeGroupPositions(s string) (res [][]int) {
	l, p, q, cnt := len(s), 0, 0, 0
	for q < l {
		if s[p] == s[q] {
			q++
			cnt++
		} else {
			if cnt >= 3 {
				res = append(res, []int{p, q - 1})
				cnt = 0
			}
			p = q
			cnt = 0
		}
	}
	if cnt >= 3 {
		res = append(res, []int{p, q - 1})
	}
	return
}

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
	fmt.Println(largeGroupPositions("abc"))
	fmt.Println(largeGroupPositions("aba"))
	fmt.Println(largeGroupPositions("aaaa"))
	fmt.Println(largeGroupPositions("aaaabbbb"))
	fmt.Println(largeGroupPositions("bababbabaa"))
}
