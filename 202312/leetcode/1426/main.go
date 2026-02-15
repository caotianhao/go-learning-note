package main

import "fmt"

func countElements1426(arr []int) (cnt int) {
	m := map[int]struct{}{}
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arr {
		if _, ok := m[v+1]; ok {
			cnt++
		}
	}
	return
}

func main() {
	fmt.Println(countElements1426([]int{1, 2, 3}))
}
