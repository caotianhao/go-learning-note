package main

import "fmt"

func getTheLongestArr(nums []int64, k int64) int {
	l := len(nums)
	prefixSum := make([]int64, l+1)
	for i := 0; i < l; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	res := -1
	for i := 0; i <= l; i++ {
		for j := i + 1; j <= l; j++ {
			if float64(prefixSum[j]-prefixSum[i])/float64(j-i) == float64(k) {
				if j-i > res {
					res = j - i
				}
			}
		}
	}
	return res
}

func main() {
	var n, k, num int64 = 0, 0, 0
	nums := make([]int64, 0)
	_, _ = fmt.Scanf("%d %d", &n, &k)
	var i int64
	for i = 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}
	fmt.Printf("%d", getTheLongestArr(nums, k))
}
