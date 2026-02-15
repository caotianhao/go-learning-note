package main

import "fmt"

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次，找出那个只出现了一次的元素
// 你的算法应该具有线性时间复杂度，你可以不使用额外空间来实现吗？
func singleNumber(nums []int) int {
	only := 0
	for i := 0; i < len(nums); i++ {
		only ^= nums[i]
	}
	return only
}

func singleNumber1(nums []int) (res int) {
	for _, v := range nums {
		res ^= v
	}
	return
}

func main() {
	fmt.Println(singleNumber([]int{2, 3, 3, 1, 4, 1, 2, 9, 9, 6, 6}))
	fmt.Println(singleNumber1([]int{2, 3, 3, 1, 4, 1, 2, 9, 9, 6, 6}))
}
