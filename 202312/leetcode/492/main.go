package main

import "fmt"

func constructRectangle(area int) []int {
	div := findAllDivisor(area)
	l := len(div)
	for i := 0; i < l; i++ {
		long, width := div[i], area/div[i]
		if long >= width {
			return []int{long, width}
		}
	}
	return []int{-1, -1}
}

func findAllDivisor(n int) (arr []int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	arr = append(arr, n)
	return
}

func main() {
	fmt.Println(constructRectangle(69))
	fmt.Println(constructRectangle(4))
}
