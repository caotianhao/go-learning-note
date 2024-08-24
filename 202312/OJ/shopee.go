package main

import "fmt"

func numSimilarGroups(strs []string) (cnt int) {
	n := len(strs)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	union := func(x, y int) {
		p[find(x)] = find(y)
	}
	isSimilar := func(a, b string) bool {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff <= 2
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}
	for i := range p {
		if p[i] == i {
			cnt++
		}
	}
	return
}

func main() {
	fmt.Println(numSimilarGroups([]string{"abcd", "cbad", "bcad", "dabc"}))
}
