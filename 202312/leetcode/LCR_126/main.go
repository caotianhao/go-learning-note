package main

import "fmt"

func fib126(n int) int {
	mod := 1000000007
	if n < 2 {
		return n
	}
	memo := make([]int, 101)
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = (memo[i-1]%mod + memo[i-2]%mod) % mod
	}
	return memo[n] % mod
}

func main() {
	fmt.Println(fib126(15))
	fmt.Println(fib126(59))
	fmt.Println(fib126(100))
}
