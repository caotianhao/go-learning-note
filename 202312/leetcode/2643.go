package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {
	maxN, index := -1, 0
	for i, v := range mat {
		cnt := 0
		for _, val := range v {
			if val == 1 {
				cnt++
			}
		}
		if cnt > maxN {
			maxN, index = cnt, i
		}
	}
	return []int{index, maxN}
}

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{{0, 1}, {1, 0}}))
}
