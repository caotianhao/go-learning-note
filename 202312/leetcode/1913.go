package main

import "fmt"

func maxProductDifference(nums []int) int {
	minN, maxN, minIndex, maxIndex, min2, max2 := nums[0], nums[0], 0, 0, 10001, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxN {
			maxN = nums[i]
			maxIndex = i
		}
		if nums[i] < minN {
			minN = nums[i]
			minIndex = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != minIndex && i != maxIndex {
			if nums[i] > max2 {
				max2 = nums[i]
			}
			if nums[i] < min2 {
				min2 = nums[i]
			}
		}
	}
	return maxN*max2 - minN*min2
}

func main() {
	arr := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(arr))
}
