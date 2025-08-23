package main

import "fmt"

func countBeautifulPairs(nums []int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if gcd001(get1(nums[i]), get0(nums[j])) == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func get1(n int) (r int) {
	for n != 0 {
		r = n % 10
		n /= 10
	}
	return r
}

func get0(n int) int {
	return n % 10
}

func gcd001(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd001(b, a%b)
}

func main() {
	fmt.Println(countBeautifulPairs([]int{2, 5, 1, 4}))
}
