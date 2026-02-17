package main

import "fmt"

func fixedPoint(arr []int) int {
	for i := range arr {
		if arr[i] == i {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9}))
}
