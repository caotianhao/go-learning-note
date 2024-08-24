package main

import "fmt"

func findLongest(s string) string {
	a, b := 0, 0
	n := len(s)
	if n == 0 {
		return ""
	}
	for i := 0; i < n; i++ {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1 > b-a {
			a, b = l1, r1
		}
		if r2-l2 > b-a {
			a, b = l2, r2
		}
	}
	return s[a : b+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func main() {
	fmt.Println(findLongest("abcbaa"))
	fmt.Println(findLongest("aaaaaaaaaaaaa"))
	fmt.Println(findLongest("aaaaaaaaaaaa"))
	fmt.Println(findLongest(""))
}
