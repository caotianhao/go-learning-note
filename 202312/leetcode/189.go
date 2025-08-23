package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse189(nums)
	reverse189(nums[:k])
	reverse189(nums[k:])
}

func reverse189(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-i-1], arr[i]
	}
}

func main() {
	a := []int{1, 2}
	fmt.Println(a)
	rotate(a, 3)
	fmt.Println(a)
}
