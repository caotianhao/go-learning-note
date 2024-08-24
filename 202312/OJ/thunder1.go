package main

import "fmt"

func thunder1(nums []int) []int {
	if len(nums) > 4 {
		return nil // return 长用例
	}
	return nil // return 短用例
}

func main() {
	fmt.Println(thunder1([]int{1, 3, 2, 5}))
}
