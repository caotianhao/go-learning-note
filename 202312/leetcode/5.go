package main

import "fmt"

func longestPalindrome(s string) string {
	l, start, end := len(s), -1, -1
	maxN := -1
	for i := 0; i < l; i++ {
		left1, right1 := expandFromCenter(i, i, &s)
		left2, right2 := expandFromCenter(i, i+1, &s)
		if right1-left1+1 > maxN {
			maxN = right1 - left1 + 1
			start, end = left1, right1
		}
		if right2-left2+1 > maxN {
			maxN = right2 - left2 + 1
			start, end = left2, right2
		}
	}
	// 这里要加一为了取全
	return s[start : end+1]
}

// 中心扩散
func expandFromCenter(left, right int, s *string) (int, int) {
	for left >= 0 && left < len(*s) && right >= 0 && right < len(*s) && (*s)[left] == (*s)[right] {
		left, right = left-1, right+1
	}
	// 扩散完出来了要缩回去
	return left + 1, right - 1
}

func main() {
	fmt.Println(longestPalindrome("ac"))
}
