package main

import "fmt"

func transformArray(arr []int) []int {
	for {
		minus, plus := make([]int, 0), make([]int, 0)
		for j := 1; j < len(arr)-1; j++ {
			if arr[j] > arr[j-1] && arr[j] > arr[j+1] {
				minus = append(minus, j)
			} else if arr[j] < arr[j-1] && arr[j] < arr[j+1] {
				plus = append(plus, j)
			}
		}
		if len(minus) == 0 && len(plus) == 0 {
			break
		}
		for _, v := range minus {
			arr[v]--
		}
		for _, v := range plus {
			arr[v]++
		}
	}
	return arr
}

func main() {
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5}))
	fmt.Println(transformArray([]int{2, 1, 2, 1, 1, 2, 2, 1}))
}
