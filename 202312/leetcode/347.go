package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) (a []int) {
	hm := make(map[int]int)
	fr := make([]int, 0)
	for _, v := range nums {
		hm[v]++
	}
	for _, v := range hm {
		fr = append(fr, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(fr)))
	f := fr[k-1]
	for key, v := range hm {
		if v >= f {
			a = append(a, key)
		}
	}
	return
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 2, 2, 2, 3}, 3))
	fmt.Println(topKFrequent([]int{1, 1, 2, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 1, 2, 2, 2, 3}, 1))
}
