package main

import "fmt"

func minOperations(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	baseIndex, needOperation, flag := 0, 0, 0
	for i := 0; i < l-1; i++ {
		if nums[i] < nums[i+1] {
			flag++
			if flag == l-1 {
				return 0
			}
			continue
		}
		baseIndex = i
		break
	}
	for i := baseIndex + 1; i < l; i++ {
		if nums[i] <= nums[i-1] {
			needOperation += nums[i-1] - nums[i] + 1
			nums[i] += nums[i-1] - nums[i] + 1
		}
	}
	return needOperation
}

func main() {
	fmt.Println(minOperations([]int{1, 5, 2, 4, 1}))
	fmt.Println(minOperations([]int{1, 5, 6, 7, 1}))
	fmt.Println(minOperations([]int{1, 5, 6, 7, 77}))
	fmt.Println(minOperations([]int{4881, 2593, 6819, 9256, 4135}))
	fmt.Println(minOperations([]int{1, 1, 1}))
}
