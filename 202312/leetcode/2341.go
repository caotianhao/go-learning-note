package main

import (
	"fmt"
)

func numberOfPairs(nums []int) []int {
	l, cntMap, cnt := len(nums), map[int]int{}, 0
	for i := 0; i < l; i++ {
		cntMap[nums[i]]++
	}
	for _, v := range cntMap {
		cnt += v / 2
	}
	return []int{cnt, l - cnt*2}
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}))
}
