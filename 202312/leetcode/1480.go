package main

import "fmt"

func runningSum(nums []int) []int {
	runSum, l := make([]int, 0), len(nums)
	for i := 0; i < l; i++ {
		tmp := 0
		for j := 0; j <= i; j++ {
			tmp += nums[j]
		}
		runSum = append(runSum, tmp)
	}
	return runSum
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4, 5}))
}
