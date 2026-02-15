package main

import "fmt"

func subsetXORSum(nums []int) int {
	l, sum := len(nums), 0
	slice1863 := make([]int, 1)
	slice1863[0] = 0
	for i := 0; i < l; i++ {
		ll := len(slice1863)
		for j := 0; j < ll; j++ {
			slice1863 = append(slice1863, nums[i]^slice1863[j])
		}
	}
	for _, v := range slice1863 {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
	fmt.Println(subsetXORSum([]int{3, 1}))
	fmt.Println(subsetXORSum([]int{3}))
}
