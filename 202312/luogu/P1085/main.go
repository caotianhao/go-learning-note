package main

import "fmt"

func main() {
	var a, b, res int
	var arr []int
	for i := 0; i < 7; i++ {
		_, _ = fmt.Scanf("%d %d\n", &a, &b)
		arr = append(arr, a+b)
	}
	maxN := -1
	for i := 0; i < 7; i++ {
		if arr[i] > maxN {
			maxN = arr[i]
		}
	}
	if maxN < 9 {
		fmt.Printf("%d", 0)
	} else {
		for i := 0; i < 7; i++ {
			if arr[i] == maxN {
				res = i + 1
				break
			}
		}
		fmt.Printf("%d", res)
	}
}
