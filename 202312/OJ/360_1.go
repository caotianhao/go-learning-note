package main

import "fmt"

func isCV(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if s == t {
		return true
	}
	if s[0] != t[0] {
		return false
	}
	slow, fast := 1, 1
	for fast < len(t) {
		if s[slow] == t[fast] {
			if slow >= len(s)-1 {
				slow = len(s) - 1
			} else {
				slow++
			}
			fast++
		} else if s[slow] != t[fast] && s[slow-1] == t[fast] {
			slow--
		} else {
			return false
		}
	}
	return true
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	var s1, s2 string
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &s1)
		_, _ = fmt.Scanf("%s", &s2)
		if isCV(s1, s2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
