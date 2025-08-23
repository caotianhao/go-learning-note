package main

import "fmt"

func maxAscendingSum(nums []int) int {
	l, maxN, sliceSum, sliceDivide := len(nums), 0, make([]int, 0), make([]int, 0)
	if l == 1 {
		return nums[0]
	}
	for i := 0; i < l-1; i++ {
		if nums[i] >= nums[i+1] {
			sliceDivide = append(sliceDivide, i)
		}
	}
	if len(sliceDivide) == 0 {
		sum := 0
		for i := 0; i < l; i++ {
			sum += nums[i]
		}
		return sum
	}
	for i := 0; i < len(sliceDivide)-1; i++ {
		sliceSum = append(sliceSum, arrSum1800(nums[sliceDivide[i]+1:sliceDivide[i+1]+1]))
	}
	sliceSum = append(sliceSum, arrSum1800(nums[:sliceDivide[0]+1]))
	sliceSum = append(sliceSum, arrSum1800(nums[sliceDivide[len(sliceDivide)-1]+1:]))
	for i := 0; i < len(sliceSum); i++ {
		if sliceSum[i] > maxN {
			maxN = sliceSum[i]
		}
	}
	return maxN
}

func arrSum1800(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	//                                 0  1   2   3  4  5   6   7   8  9
	//fmt.Println(maxAscendingSum([]int{10, 1, 20, 30, 5, 10, 50, 10, 3, 1})) //65
	//fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50})) //65
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 40, 50})) //150
	//fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12})) //33
	//fmt.Println(maxAscendingSum([]int{100, 10, 1})) //100
	//fmt.Println(maxAscendingSum([]int{333})) //333
}
