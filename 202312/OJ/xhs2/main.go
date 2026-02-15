package main

import (
	"fmt"
	"strings"
)

// w = vv
// m = nn
// b == d q p
// p == q d b
// d == p q b
// q == b p d
// n == u

func judge0(a, c byte) bool {
	if a == c {
		return true
	}
	if a == 'b' && (c == 'd' || c == 'p' || c == 'q') {
		return true
	}
	if a == 'd' && (c == 'b' || c == 'p' || c == 'q') {
		return true
	}
	if a == 'p' && (c == 'd' || c == 'b' || c == 'q') {
		return true
	}
	if a == 'q' && (c == 'd' || c == 'p' || c == 'b') {
		return true
	}
	if a == 'u' && c == 'n' {
		return true
	}
	if a == 'n' && c == 'u' {
		return true
	}
	return false
}

func judge(s string) bool {
	haveW := false
	haveM := false
	n := len(s)
	r := []byte(s)
	for _, v := range s {
		if v == 'w' {
			haveW = true
			break
		}
	}
	for _, v := range s {
		if v == 'm' {
			haveM = true
			break
		}
	}
	if !haveW {
		if !haveM {
			for i := 0; i < n/2; i++ {
				if !judge0(r[i], r[n-1-i]) {
					return false
				}
			}
			return true
		} else {
			s = strings.ReplaceAll(s, "m", "nn")
			return judge(s)
		}
	} else {
		if !haveM {
			s = strings.ReplaceAll(s, "w", "vv")
			return judge(s)
		} else {
			s = strings.ReplaceAll(s, "m", "nn")
			s = strings.ReplaceAll(s, "w", "vv")
			return judge(s)
		}
	}
}

func main() {
	n := 0
	s := ""
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &s)
		if judge(s) {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}
}
