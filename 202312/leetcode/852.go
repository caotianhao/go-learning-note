package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l := len(arr)
	for i := 1; i < l-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
}
