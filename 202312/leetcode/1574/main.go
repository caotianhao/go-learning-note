package main

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	res := j
	for i := 0; i < n; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		res = min(res, j-i-1)
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})) //3
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))           //4
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))                 //0
	fmt.Println(findLengthOfShortestSubarray([]int{1}))                       //0
}
