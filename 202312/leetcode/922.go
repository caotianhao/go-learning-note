package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	slice922, ji, ou := make([]int, len(nums)), 1, 0
	for _, v := range nums {
		if v%2 == 0 {
			slice922[ou] = v
			ou += 2
		} else {
			slice922[ji] = v
			ji += 2
		}
	}
	return slice922
}

func main() {
	fmt.Println(sortArrayByParityII([]int{1, 2, 3, 4}))
}
