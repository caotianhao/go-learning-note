package main

import "fmt"

func balancedStringSplit(s string) int {
	l, ll, rr, cnt := len(s), 0, 0, 0
	for i := 0; i < l; i++ {
		if s[i] == 'L' {
			ll++
		}
		if s[i] == 'R' {
			rr++
		}
		if ll == rr {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(balancedStringSplit("LRRLRRLLRL"))
}
