package main

import (
	"fmt"
	"sort"
)

func maxSum(nums []int) int {
	dig := make([]int, len(nums))
	for i, v := range nums {
		dig[i] = getMaxDigit(v)
	}
	maxN := -1
	for i := 0; i < len(dig); i++ {
		for j := i + 1; j < len(dig); j++ {
			if dig[i] == dig[j] {
				t := nums[i] + nums[j]
				if t > maxN {
					maxN = t
				}
			}
		}
	}
	return maxN
}

func getMaxDigit(n int) int {
	digit := make([]int, 0)
	for n != 0 {
		digit = append(digit, n%10)
		n /= 10
	}
	sort.Ints(digit)
	return digit[len(digit)-1]
}

func main() {
	fmt.Println(maxSum([]int{51, 71, 17, 24, 42}))
}
