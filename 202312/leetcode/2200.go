package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) (res []int) {
	id, l := make([]int, 0), len(nums)
	for i, v := range nums {
		if v == key {
			id = append(id, i)
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < len(id); j++ {
			if abs2200(i, id[j]) <= k {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func abs2200(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(findKDistantIndices([]int{1, 2, 3, 4}, 1, 2))
}
