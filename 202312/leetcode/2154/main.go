package main

import "fmt"

func findFinalValue(nums []int, original int) int {
	if !isNumInArr2154(nums, original) {
		return original
	} else {
		return findFinalValue(nums, original*2)
	}
}

func isNumInArr2154(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
	fmt.Println(findFinalValue([]int{2, 7, 9}, 4))
	fmt.Println(findFinalValue([]int{2, 7, 9}, 0))
	fmt.Println(findFinalValue([]int{0, 7, 9}, 0))
}
