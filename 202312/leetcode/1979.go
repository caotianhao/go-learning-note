package main

import (
	"fmt"
	"sort"
)

func findGCD1(nums []int) int {
	minN, maxN := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minN {
			minN = nums[i]
		}
		if nums[i] > maxN {
			maxN = nums[i]
		}
	}
	return gcd1979(maxN, minN)
}

func gcd1979(i int, j int) int {
	if j > i {
		i, j = j, i
	}
	for i%j != 0 {
		return gcd1979(j, i%j)
	}
	return j
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	var gcd func(int, int) int
	gcd = func(a int, b int) int {
		if b != 0 {
			return gcd(b, a%b)
		}
		return a
	}
	return gcd(nums[0], nums[len(nums)-1])
}

func main() {
	a := []int{8, 46, 512, 1024, 5508}
	fmt.Println(findGCD1(a))
	fmt.Println(findGCD(a))
}
