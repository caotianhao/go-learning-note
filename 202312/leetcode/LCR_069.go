package main

import "fmt"

func peakIndexInMountainArray2069(arr []int) int {
	l := len(arr)
	for i := 1; i < l-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray2069([]int{1, 2, 6, 4, 8}))
}
