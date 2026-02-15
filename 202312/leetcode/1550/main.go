package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	l := len(arr)
	if l < 3 {
		return false
	}
	for i := 0; i < l-2; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{1, 2}))
}
