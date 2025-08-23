package main

import "fmt"

func specialArray(nums []int) int {
	var cnt, special int
	for special = 0; special <= 100; special++ {
		cnt = 0
		for i := 0; i < len(nums); i++ {
			if special <= nums[i] {
				cnt++
			}
		}
		if special == cnt {
			return special
		}
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{0, 0, 3, 4, 4}))
}
