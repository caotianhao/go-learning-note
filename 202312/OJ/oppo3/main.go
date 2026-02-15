package main

import "fmt"

func minKP(nums []int, k int) (int, int) {
	product := 1
	for _, v := range nums {
		product *= v
	}
	a1, a2 := 0, 0
	now := countFactor(product)
	if now == k {
		return a1, a2
	}
	return 12, 0
}

func countFactor(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	n := 0
	k := 0
	a := 0
	arr := make([]int, 0)
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	aa, bb := minKP(arr, k)
	fmt.Printf("%d %d", aa, bb)
}
