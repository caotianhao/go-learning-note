package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	l, sum := len(arr), 0
	for i := 1; i <= l; i += 2 {
		for j := 0; j < l+1 && j+i < l+1; j++ {
			sum += arrSum(arr[j : j+i])
		}
	}
	return sum
}

func arrSum(arr []int) int {
	l, sum := len(arr), 0
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
