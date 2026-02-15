package main

import (
	"fmt"
	"strings"
)

func countAB(s string, n, k int) int {
	//i := 1
	//for ; ; i++ {
	//	if s[(i-1)%n] == 'A' {
	//		k--
	//	}
	//	if k == 0 {
	//		break
	//	}
	//}
	//return i
	idx := make([]int, 0)
	cntA := strings.Count(s, "A")
	for i, v := range s {
		if v == 'A' {
			idx = append(idx, i)
		}
	}
	ka, kb := k/cntA, k%cntA
	if kb == 0 {
		return n * ka
	}
	return n*ka + 1 + idx[kb-1]
}

func main() {
	n, k := 0, 0
	var ab string
	_, _ = fmt.Scanf("%d %d", &n, &k)
	_, _ = fmt.Scanf("%s", &ab)
	fmt.Printf("%d", countAB(ab, n, k))
}
