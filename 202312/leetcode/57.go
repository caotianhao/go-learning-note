package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	idx := sort.Search(len(intervals), func(temp int) bool {
		return intervals[temp][0] >= newInterval[0]
	})
	queue := make([][]int, 0)
	for i := 0; i < idx; i++ {
		queue = append(queue, intervals[i])
	}
	queue = append(queue, newInterval)
	for i := idx; i < len(intervals); i++ {
		queue = append(queue, intervals[i])
	}
	res = append(res, queue[0])
	for _, v := range queue[1:] {
		if v[0] > res[len(res)-1][1] {
			res = append(res, v)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], v[1])
		}
	}
	return
}

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
