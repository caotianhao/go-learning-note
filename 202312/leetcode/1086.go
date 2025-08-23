package main

import (
	"fmt"
	"sort"
)

func highFive(items [][]int) (res [][]int) {
	m := map[int][]int{}
	for _, v := range items {
		m[v[0]] = append(m[v[0]], v[1])
	}
	for i, v := range m {
		if len(v) != 0 {
			sort.Sort(sort.Reverse(sort.IntSlice(v)))
			res = append(res, []int{i, (v[0] + v[1] + v[2] + v[3] + v[4]) / 5})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return
}

func main() {
	fmt.Println(highFive([][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77},
		{1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}}))
}
