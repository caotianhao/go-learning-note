package main

import (
	"fmt"
	"sort"
)

func maxCount(m int, n int, ops [][]int) int {
	oo := deleteZero598(ops)
	if len(oo) == 0 {
		return m * n
	}
	zero, one := make([]int, 0), make([]int, 0)
	for _, v := range oo {
		zero = append(zero, v[0])
		one = append(one, v[1])
	}
	sort.Ints(zero)
	sort.Ints(one)
	return zero[0] * one[0]
}

func deleteZero598(arr [][]int) (ret [][]int) {
	for _, v := range arr {
		if !(v[0] == 0 || v[1] == 0) {
			ret = append(ret, v)
		}
	}
	return
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	fmt.Println(deleteZero598([][]int{{2, 2}, {3, 3}, {1, 0}, {2, 0}, {2, 3}, {0, 2}, {2, 6}}))
}
