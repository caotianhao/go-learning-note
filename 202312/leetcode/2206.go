package main

import "fmt"

func divideArray(nums []int) bool {
	l, cntMap := len(nums), map[int]int{}
	for i := 0; i < l; i++ {
		cntMap[nums[i]]++
	}
	for _, v := range cntMap {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(divideArray([]int{2, 2, 3, 3}))
}
