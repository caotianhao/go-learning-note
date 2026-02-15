package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	cnt := 0
	sort.Ints(g)
	sort.Ints(s)
	need, have := 0, 0
	for have < len(s) && need < len(g) {
		if g[need] <= s[have] {
			cnt++
			need++
			have++
		} else {
			have++
		}
	}
	return cnt
}

func main() {
	fmt.Println(findContentChildren([]int{1, 1, 2}, []int{1, 2, 2}))
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{7, 8, 11, 13}, []int{5, 6, 9, 10}))
}
