package main

import (
	"fmt"
	"sort"
)

func check(nums []int) bool {
	source := make([]int, 0)
	for _, v := range nums {
		source = append(source, v)
	}
	sort.Ints(source)
	l := len(nums)
	for i := 0; i < l; i++ {
		cnt := 0
		for j := 0; j < l; j++ {
			if source[j] == nums[(j+i)%l] {
				cnt++
			}
		}
		if cnt == l {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 2, 1}))
}
