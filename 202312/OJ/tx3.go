package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	a, arr := 0, make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	// 分两段
	sort.Ints(arr)
	l := len(arr)
	var x, d, fin []int
	x = arr[:l/2]
	d = arr[l/2:]
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	lx, ld := len(x), len(d)
	for i := 0; i < min(lx, ld); i++ {
		fin = append(fin, d[i])
		fin = append(fin, x[i])
	}
	if lx > ld {
		for i := 0; i < lx-ld; i++ {
			fin = append(fin, x[i+ld])
		}
	} else {
		for i := 0; i < ld-lx; i++ {
			fin = append(fin, d[i+lx])
		}
	}
	force := 0
	final := make([]int, 1)
	final = append(final, fin...)
	//fmt.Println(x, d, final)
	//if len(final ) % 2==0 {
	for i := 0; i < len(final); i += 2 {
		if i+1 == len(final) {
			//force += final[(len(final) - 1)]
			break
		}
		force += final[i+1] - final[i]
	}
	//} else {
	//
	//}
	fmt.Printf("%d", force)
}
