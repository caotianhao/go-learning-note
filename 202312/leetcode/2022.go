package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	my2Arr := make([][]int, 0)
	if len(original) != m*n {
		return my2Arr
	}
	for i := 0; i < m; i++ {
		mySlice := make([]int, 0)
		for j := 0; j < n; j++ {
			mySlice = append(mySlice, original[j+i*len(original)/m])
		}
		my2Arr = append(my2Arr, mySlice)
	}
	return my2Arr
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(construct2DArray(arr, 2, 2))
}
