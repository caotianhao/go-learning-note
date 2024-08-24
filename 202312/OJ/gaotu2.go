package main

import "fmt"

func lengthSubString(s string) int {
	hm := make(map[byte]int)
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(hm, s[i-1])
		}
		for right+1 < n && hm[s[right+1]] == 0 {
			hm[s[right+1]]++
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	fmt.Println(lengthSubString(s))
}
