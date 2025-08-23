package main

import (
	"fmt"
	"math/bits"
)

func sumIndicesWithKSetBits(nums []int, k int) (sum int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		cnt := bits.OnesCount64(uint64(i))
		if cnt == k {
			sum += nums[i]
		}
	}
	return
}

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
}
