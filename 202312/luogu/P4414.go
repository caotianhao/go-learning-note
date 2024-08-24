package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int
	var arr []int
	var c1, c2, c3 byte
	_, _ = fmt.Scanf("%d %d %d\n", &n1, &n2, &n3)
	arr = append(arr, n1, n2, n3)
	sort.Ints(arr)
	a, b, c := arr[0], arr[1], arr[2]
	_, _ = fmt.Scanf("%c%c%c", &c1, &c2, &c3)
	if c1 == 'A' && c2 == 'B' && c3 == 'C' {
		fmt.Printf("%d %d %d", a, b, c)
	} else if c1 == 'A' && c2 == 'C' && c3 == 'B' {
		fmt.Printf("%d %d %d", a, c, b)
	} else if c1 == 'B' && c2 == 'A' && c3 == 'C' {
		fmt.Printf("%d %d %d", b, a, c)
	} else if c1 == 'B' && c2 == 'C' && c3 == 'A' {
		fmt.Printf("%d %d %d", b, c, a)
	} else if c1 == 'C' && c2 == 'A' && c3 == 'B' {
		fmt.Printf("%d %d %d", c, a, b)
	} else {
		fmt.Printf("%d %d %d", c, b, a)
	}
}
