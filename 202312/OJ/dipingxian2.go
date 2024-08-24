package main

import "fmt"

func maxArea(height []int) (ans int) {
	n := len(height)
	left, right := 0, n-1
	for left < right {
		tmpArea := min(height[left], height[right]) * (right - left)
		ans = max(ans, tmpArea)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func main() {
	fmt.Println(maxArea([]int{1, 7, 3, 2, 4, 5, 8, 2, 7})) //49
}
