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
	fmt.Printf("%d", arr[2])
}
