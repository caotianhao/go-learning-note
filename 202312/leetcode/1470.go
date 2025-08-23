package main

import "fmt"

func shuffle(nums []int, n int) []int {
	part1, part2, all, flag := make([]int, 0), make([]int, 0), make([]int, 2*n), 0
	for i := 0; i < n; i++ {
		part1 = append(part1, nums[i])
	}
	for i := n; i < 2*n; i++ {
		part2 = append(part2, nums[i])
	}
	for i := 0; i < n; i++ {
		if flag%2 == 0 {
			all[flag] = part1[i]
			flag++
		}
		if flag%2 == 1 {
			all[flag] = part2[i]
			flag++
		}
	}
	return all
}

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	fmt.Println(shuffle(arr, 3))
}
