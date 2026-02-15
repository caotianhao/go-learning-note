package main

import (
	"fmt"
	"math"
)

func maxLengthWithCombineAB(a, b string) int {
	queue := getPermutation(a, b)
	res := math.MinInt64
	for _, s := range queue {
		res = max(res, longestP(s))
	}
	return res
}

func longestP(s string) int {
	n := len(s)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		l1, r1 := expandFromCenter(s, i, i)
		l2, r2 := expandFromCenter(s, i, i+1)
		if r1-l1 > r-l {
			l, r = l1, r1
		}
		if r2-l2 > r-l {
			l, r = l2, r2
		}
	}
	return r + 1 - l
}

func expandFromCenter(s string, l, r int) (int, int) {
	n := len(s)
	for ; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}

func getPermutation(a, b string) (ans []string) {
	var dfs func(string, string, string)
	dfs = func(path string, remainingA string, remainingB string) {
		if remainingA == "" && remainingB == "" {
			ans = append(ans, path)
			return
		}
		if remainingA != "" {
			dfs(path+string(remainingA[0]), remainingA[1:], remainingB)
		}
		if remainingB != "" {
			dfs(path+string(remainingB[0]), remainingA, remainingB[1:])
		}
	}
	dfs("", a, b)
	return
}

func main() {
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	var a, b string
	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%s", &a)
		_, _ = fmt.Scanf("%s", &b)
		fmt.Println(maxLengthWithCombineAB(a, b))
	}
}
