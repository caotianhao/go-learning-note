package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	slice905 := make([]int, 0)
	for _, v := range nums {
		if v%2 == 0 {
			slice905 = append(slice905, v)
		}
	}
	for _, v := range nums {
		if v%2 == 1 {
			slice905 = append(slice905, v)
		}
	}
	return slice905
}

func main() {
	fmt.Println(sortArrayByParity([]int{1, 2, 3, 4, 5, 6, 8, 10}))
}
