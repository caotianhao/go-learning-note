package main

import "fmt"

func decompressRLElist(nums []int) []int {
	myReturn, l := make([]int, 0), len(nums)
	for i := 0; i < l; i += 2 {
		for j := 0; j < nums[i]; j++ {
			myReturn = append(myReturn, nums[i+1])
		}
	}
	return myReturn
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(arr))
}
