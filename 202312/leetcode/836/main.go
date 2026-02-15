package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	b1 := min(rec1[2], rec2[2]) > max(rec1[0], rec2[0])
	b2 := min(rec1[3], rec2[3]) > max(rec1[1], rec2[1])
	return b1 && b2
}

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
}
