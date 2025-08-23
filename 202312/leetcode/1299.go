package main

import "fmt"

func replaceElements(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		if i != l-1 {
			arr[i] = 0
			for j := i + 1; j < l; j++ {
				if arr[j] > arr[i] {
					arr[i] = arr[j]
				}
			}
		} else {
			arr[l-1] = -1
		}
	}
	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
