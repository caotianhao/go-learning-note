package main

import (
	"fmt"
)

func minValueAfterModification(s string, k int) int {
	p := make([][2]int, 0)
	tmp := s[0]
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == tmp {
			count++
		} else {
			p = append(p, [2]int{int(tmp - '0'), count})
			tmp = s[i]
			count = 1
		}
	}

	p = append(p, [2]int{int(tmp - '0'), count})

	if k >= len(p) {
		return len(s)
	} else {
		maxValue := 0
		for i := 0; i <= k; i++ {
			maxValue += p[i][1]
		}
		return maxValue
	}
}

func main() {
	s := "10101100011011010"
	k := 3
	result := minValueAfterModification(s, k)
	fmt.Println(result) // 输出 11
}
