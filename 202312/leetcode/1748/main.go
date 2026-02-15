package main

import "fmt"

func sumOfUnique(nums []int) int {
	cnt, l, sum := make([]int, 101), len(nums), 0
	for i := 0; i < l; i++ {
		cnt[nums[i]]++
	}
	fmt.Println(cnt)
	for i := 0; i < 101; i++ {
		if cnt[i] == 1 {
			sum += i
		}
	}
	return sum
}

func main() {
	arr := []int{1, 2, 2, 9, 2, 2, 2, 2, 2, 2}
	fmt.Println(sumOfUnique(arr))
}
