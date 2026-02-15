package main

import (
	"fmt"
	"sort"
)

func main() {
	q := 0
	_, _ = fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		n, m := 0, 0
		_, _ = fmt.Scanf("%d %d", &n, &m)
		aa, bb := make([]int, 0), make([]int, 0)
		a, b := 0, 0
		for j := 0; j < n; j++ {
			_, _ = fmt.Scanf("%d", &a)
			aa = append(aa, a)
		}
		for j := 0; j < n; j++ {
			_, _ = fmt.Scanf("%d", &b)
			bb = append(bb, b)
		}
		// ------------------------------
		sort.Sort(sort.Reverse(sort.IntSlice(aa)))
		sort.Ints(bb)
		flag := true
		for window := 0; window < n; window++ {
			if aa[window]+bb[window] > m {
				flag = false
			}
		}
		if flag {
			fmt.Printf("%s\n", "Yes")
		} else {
			fmt.Printf("%s\n", "No")
		}
	}
}
