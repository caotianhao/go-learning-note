package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	r, ret := len(score), make([][]int, 0)
	kthScore, kthScoreRow := make([]int, 0), make([]int, 0)
	for i := 0; i < r; i++ {
		kthScore = append(kthScore, score[i][k])
	}
	sort.Slice(kthScore, func(i, j int) bool {
		return kthScore[i] > kthScore[j]
	})
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if kthScore[i] == score[j][k] {
				kthScoreRow = append(kthScoreRow, j)
			}
		}
	}
	for i := 0; i < r; i++ {
		ret = append(ret, score[kthScoreRow[i]])
	}
	return ret
}

func main() {
	fmt.Println(sortTheStudents([][]int{
		{10, 6, 9, 1},
		{7, 5, 11, 2},
		{4, 8, 3, 15},
	}, 1))
	fmt.Println(sortTheStudents([][]int{
		{3, 4},
		{5, 6},
	}, 0))
}
