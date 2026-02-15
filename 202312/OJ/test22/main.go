package main

import (
	"fmt"
	"math"
	"sort"
)

func maxRightShift(n int64) int64 {
	l := 0
	t := n
	for t != 0 {
		t /= 10
		l++
	}
	ans := make([]int64, 0)
	for i := 0; i < l; i++ {
		n = rs(n, l)
		ans = append(ans, n)
	}
	fmt.Println(ans)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	fmt.Println(ans)
	return ans[0]
}

func rs(n int64, l int) int64 {
	k := n % 10
	k2 := n / 10
	return k*int64(math.Pow(10.0, float64(l-1))) + k2
}

func main() {
	fmt.Println(math.MaxInt64 + 3000000000 - 20000000000)
	fmt.Println(maxRightShift(math.MaxInt64 + 3000000000 - 20000000000))
}
