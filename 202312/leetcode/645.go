package main

import "fmt"

func findErrorNums(nums []int) []int {
	l, two, lost := len(nums), -1, -1
	for i := 0; i < l; {
		if nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				two = nums[i]
				i++
			} else {
				nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			}
		} else {
			i++
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			lost = i + 1
			break
		}
	}
	return []int{two, lost}
}

func main() {
	fmt.Println(findErrorNums([]int{4, 2, 2, 1}))
	fmt.Println(findErrorNums([]int{9, 3, 6, 5, 6, 1, 7, 4, 2}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{2, 2}))
	fmt.Println(findErrorNums([]int{1, 2, 13, 10, 4, 5, 6, 11, 8, 9, 7, 12, 10}))
}
