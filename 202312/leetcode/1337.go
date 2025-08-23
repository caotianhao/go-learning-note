package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	l, army, tmp := len(mat), map[int]int{}, make([]int, 0)
	for i := 0; i < l; i++ {
		army[i] = howManyOne1337(mat[i])
		tmp = append(tmp, i)
	}
	sort.Slice(tmp, func(i, j int) bool {
		return (army[tmp[i]] < army[tmp[j]]) ||
			(army[tmp[i]] == army[tmp[j]] && tmp[i] < tmp[j])
	})
	return tmp[:k]
}

func howManyOne1337(arr []int) int {
	cnt := 0
	for _, v := range arr {
		if v == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 0))
	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, 2))
}
