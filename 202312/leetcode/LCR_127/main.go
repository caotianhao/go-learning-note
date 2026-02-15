package main

import "fmt"

func trainWays(num int) int {
	mod := 1000000007
	memo := make([]int, 101)
	memo[1], memo[0] = 1, 1
	for i := 2; i <= num; i++ {
		memo[i] = (memo[i-1]%mod + memo[i-2]%mod) % mod
	}
	return memo[num] % mod
}

func main() {
	fmt.Println(trainWays(15))
	fmt.Println(trainWays(2))
	fmt.Println(trainWays(5))
}
