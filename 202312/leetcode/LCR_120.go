package main

import "fmt"

func findRepeatDocument(documents []int) int {
	hm := make(map[int]int)
	for _, document := range documents {
		hm[document]++
		if hm[document] > 1 {
			return document
		}
	}
	return -1
}

func main() {
	fmt.Println(findRepeatDocument([]int{1, 2, 3, 3}))
}
