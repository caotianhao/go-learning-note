package main

import "fmt"

// 方法 1
func findDuplicates1(nums []int) (ans []int) {
	// 题意说了数组中的元素范围在 [1, n]
	// 正好减去 1 之后对应下标
	// 由于不能使用额外空间，再结合题意
	// 使用正负号标记是否出现过
	for i := range nums {
		// 由于在遍历到该位置之前可能被加了负号
		// 所以这里要取绝对值
		t := abs442(nums[i])
		// 先判断如果位置处为小于 0 的数，就说明该位置代表数字之前出现过
		if nums[t-1] < 0 {
			ans = append(ans, t)
		}
		// 下标正负标记法
		nums[t-1] *= -1
	}
	return
}

func abs442(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// 方法 2
func findDuplicates2(nums []int) (ans []int) {
	// 例如 nums[0] = 5，则把其与下标 4 的数交换
	// 直到不能再交换
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// 此时对应位置数字错误的即为答案
	for i, v := range nums {
		if i+1 != v {
			ans = append(ans, v)
		}
	}
	return
}

func main() {
	fmt.Println(findDuplicates1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates1([]int{2, 2, 3}))
	fmt.Println(findDuplicates2([]int{1}))
	fmt.Println(findDuplicates2([]int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6}))
	//fmt.Println(4 & 3)
}
