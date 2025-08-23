package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	l, minN := len(nums), 100001
	sort.Ints(nums)
	if l == k {
		return nums[l-1] - nums[0]
	}
	if k == 1 {
		return 0
	}
	for i := 0; i <= l-k; i++ {
		if nums[i+k-1]-nums[i] < minN {
			minN = nums[i+k-1] - nums[i]
		}
	}
	return minN
}

func main() {
	arr := []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
	fmt.Println(minimumDifference(arr, 6))
	sort.Ints(arr)
	fmt.Println(arr)
}
