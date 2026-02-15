package main

import (
	"fmt"
	"math"
)

func maxSumDivThree(nums []int) int {
	//l := len(nums)
	//dp := make([][]int, l+1)
	dp := []int{0, math.MinInt64, math.MinInt64}
	//for i := 0; i < l+1; i++ {
	//	dp[i] = make([]int, 3)
	//}
	//dp[1], dp[2] = math.MinInt64, math.MinInt64
	//dp = append(dp, 0)
	for i := 1; i < 1+len(nums); i++ {
		//for j := 0; j < 3; j++ {
		//	dp[i][(j+nums[i-1])%3] = max1262(dp[i-1][(j+nums[i-1])%3], dp[i-1][j]+nums[i-1])
		//}

		//zero, one, two := make([]int, 0), make([]int, 0), make([]int, 0)
		//a, b, c := dp[0], dp[1], dp[2]
		//d, e, f := a+nums[i-1], b+nums[i-1], c+nums[i-1]
		//for _, v := range []int{a, b, c, d, e, f} {
		//	switch v % 3 {
		//	case 0:
		//		zero = append(zero, v)
		//	case 1:
		//		one = append(one, v)
		//	case 2:
		//		two = append(two, v)
		//	}
		//}
		//if len(zero) != 0 {
		//	dp[0] = max1262(zero)
		//} else {
		//	dp[0] = a
		//}
		//if len(one) != 0 {
		//	dp[1] = max1262(one)
		//} else {
		//	dp[1] = b
		//}
		//if len(two) != 0 {
		//	dp[2] = max1262(two)
		//} else {
		//	dp[2] = c
		//}
		ddp := make([]int, 3)
		for j := 0; j < 3; j++ {
			// 节省空间的做法
			// 不选当前数，则为上一行的值
			// 选了就是反向计算余数，加入进来
			//ddp[(j+nums[i-1])%3] = max1262(dp[(j+nums[i-1])%3], dp[j]+nums[i-1])
			ddp[j] = max(dp[j], dp[getMod(j-nums[i-1], 3)]+nums[i-1])
		}
		dp = ddp

	}
	return dp[0]
}

//func max1262(a []int) int {
//	sort.Ints(a)
//	return a[len(a)-1]
//}

func getMod(a, b int) int {
	r := a % b
	for r < 0 {
		r += b
	}
	return r
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
}
