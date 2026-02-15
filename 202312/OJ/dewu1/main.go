package main

import (
	"fmt"
	"sort"
)

func height(h []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(h)))
	n := len(h)
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			right = append(right, h[i])
		} else {
			left = append(left, h[i])
		}
	}
	left = reverse(left)
	left = append(left, h[0])
	for i := 0; i < len(right); i++ {
		left = append(left, right[i])
	}
	res := -1
	fmt.Println(left)
	for i := 0; i < len(left)-1; i++ {
		t := abs(left[i] - left[i+1])
		if t > res {
			res = t
		}
	}
	return res
}

func reverse(a []int) []int {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	n, h := 0, 0
	hs := make([]int, 0)
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &h)
		hs = append(hs, h)
	}
	fmt.Printf("%d", height(hs))
}
