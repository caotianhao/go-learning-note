package main

import (
	"fmt"
	"math/rand"
	"time"
)

// i < j < k
// nums[j] - nums[i] == diff
// nums[k] - nums[j] == diff
func arithmeticTriplets(nums []int, diff int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	arr := make([]int, 0)
	for i := 0; i < 300; i++ {
		arr = append(arr, rand.Intn(200)+1)
	}
	time0 := time.Now()
	a := arithmeticTriplets(arr, 3)
	time1 := time.Since(time0)
	fmt.Println(a)
	fmt.Println(time1)
}
