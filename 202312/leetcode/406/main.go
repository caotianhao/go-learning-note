package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] ||
			people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	for _, p := range people {
		idx := p[1]
		ans = append(ans[:idx], append([][]int{p}, ans[idx:]...)...)
	}
	return
}

func main() {
	// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
