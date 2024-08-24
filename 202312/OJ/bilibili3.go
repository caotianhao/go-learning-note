package main

import "fmt"

func swapByMod(nums []int, k int) []int {
	for i := 1; i <= k; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j]%i > nums[j+1]%i {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	n, m := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &m)
	a, arr := 0, make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	ans := swapByMod(arr, m)
	for _, v := range ans {
		fmt.Printf("%d\n", v)
	}
}
