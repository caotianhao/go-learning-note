package main

import (
	"fmt"
	"sort"
)

//func sumOfPower(nums []int) int {
//	sort.Ints(nums)
//	l := len(nums)
//	var s int
//	dp := make([][]int, l)
//	for i := 0; i < l; i++ {
//		dp[i] = make([]int, l)
//	}
//	for i := 0; i < l; i++ {
//		for j := i; j < l; j++ {
//			if i == j {
//				dp[i][j] = ((nums[j] * nums[j] % mod) * (nums[i] % mod)) % mod
//			} else {
//				mi := 1
//				for k := 0; k < j-i-1; k++ {
//					mi = mi * 2 % mod
//				}
//				dp[i][j] = ((nums[j] * nums[j]) % mod) * (nums[i] * mi % mod) % mod
//			}
//		}
//	}
//	for i := 0; i < l; i++ {
//		for j := i; j < l; j++ {
//			s = (s + dp[i][j]) % mod
//		}
//	}
//	return s
//}

//func sumOfPower(nums []int) int {
//	sort.Ints(nums)
//	ans := 0
//	for i := 0; i < len(nums); i++ {
//		ans = (ans + nums[i]*nums[i]*nums[i]) % mod
//		for j := i + 1; j < len(nums); j++ {
//			ans = (ans + (1<<(j-i-1))%mod*nums[j]%mod*nums[j]) % mod
//		}
//	}
//	return ans
//}

//func sumOfPower(nums []int) int {
//	n := len(nums)
//	sort.Ints(nums)
//	dp := make([]int, n)
//	preSum := make([]int, n+1)
//	res := 0
//	for i := 0; i < n; i++ {
//		dp[i] = (nums[i] + preSum[i]) % mod
//		preSum[i+1] = (preSum[i] + dp[i]) % mod
//		res = (res + (nums[i]*nums[i]%mod*dp[i])%mod) % mod
//	}
//	fmt.Println(dp)
//	fmt.Println(preSum)
//	return res
//}

func sumOfPower(nums []int) (ans int) {
	sort.Ints(nums)
	temp, mod := 0, int(1e9+7)
	for _, d := range nums {
		ans = (ans + d*d%mod*(d+temp)) % mod
		temp = (2*temp + d) % mod
	}
	return
}

func main() {
	fmt.Println(sumOfPower([]int{2, 1, 4}))
	fmt.Println(sumOfPower([]int{658, 489, 777, 2418, 1893, 130, 2448, 178, 1128,
		2149, 1059, 1495, 1166, 608, 2006, 713, 1906, 2108, 680, 1348, 860, 1620,
		146, 2447, 1895, 1083, 1465, 2351, 1359, 1187, 906, 533, 1943, 1814, 1808,
		2065, 1744, 254, 1988, 1889, 1206}))
	fmt.Println(sumOfPower([]int{1705, 7581, 5816, 8226, 3173, 2381, 9562, 5447,
		8770, 2247, 2106, 6925, 4347, 4651, 3785, 7013, 4826, 1260, 9151, 5321,
		6521, 2984, 2553, 6035, 4095, 9021, 8296, 7682, 4071, 2830, 4182, 1994,
		9222, 5343, 6826, 330, 8214, 5657, 425, 7030, 9074, 4016, 1393, 4598, 9563,
		2811, 9156, 9399, 140, 5627, 6590, 9980, 6620, 2461, 6213, 3966, 301, 6254,
		8024, 3508, 2979, 1707, 4293, 3789, 7349, 728, 5802, 4223, 2195, 9914, 2294,
		2551, 8221, 3807, 3860, 4967, 2506, 3086, 7531, 6922, 8546, 5636, 6249}))
}
