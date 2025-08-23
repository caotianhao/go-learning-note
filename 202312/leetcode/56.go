package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) (res [][]int) {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < n; i++ {
		if res[len(res)-1][0] <= intervals[i][0] && intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		}
	}
	return
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 9}, {8, 10}, {19, 33}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}}))
}
