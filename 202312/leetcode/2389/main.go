package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	ln, lq, slice2389 := len(nums), len(queries), make([]int, 0)
	sort.Ints(nums)
	for i := 0; i < lq; i++ {
		sum, index := 0, -1
		for j := 0; j < ln; j++ {
			sum += nums[j]
			if sum > queries[i] {
				index = j
				break
			}
			if j == ln-1 {
				index = ln
			}
		}
		slice2389 = append(slice2389, index)
	}
	return slice2389
}

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
	fmt.Println(answerQueries([]int{7, 9, 1, 2, 3, 6}, []int{1, 2, 6, 9, 22, 300}))
	fmt.Println(answerQueries([]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5, 6}))
}
