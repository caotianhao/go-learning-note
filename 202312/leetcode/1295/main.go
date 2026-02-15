package main

import "fmt"

//1 <= nums[i] <= 100000
func findNumbers(nums []int) int {
	cnt, l := 0, len(nums)
	for i := 0; i < l; i++ {
		if nums[i] >= 10 && nums[i] <= 99 {
			cnt++
		}
		if nums[i] >= 1000 && nums[i] <= 9999 {
			cnt++
		}
		if nums[i] == 100000 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
