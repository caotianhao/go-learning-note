package main

import (
	"fmt"
	"sort"
)

type help1338 struct {
	key   int
	value int
}

func minSetSize(arr []int) (c int) {
	l, t := len(arr), 0
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	h := make([]help1338, 0)
	for k, v := range m {
		h = append(h, help1338{key: k, value: v})
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].value > h[j].value
	})
	for i := range h {
		t += h[i].value
		c++
		if t >= l/2 {
			return
		}
	}
	return
}

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
