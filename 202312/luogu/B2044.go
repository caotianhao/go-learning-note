package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	arr := []int{a, b, c}
	sort.Ints(arr)
	if a < 60 && b >= 60 && c >= 60 {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}
}
