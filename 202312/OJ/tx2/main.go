package main

import (
	"fmt"
	"sort"
)

func popA(arr *[]int, n int) (r []int) {
	k := 0
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == n {
			k++
		}
	}
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] != n {
			r = append(r, (*arr)[i])
		}
	}
	for i := 0; i < k-1; i++ {
		r = append(r, n)
	}
	return
}

func mid(arr *[]int) string {
	sort.Ints(*arr)
	l := len(*arr)
	if l%2 == 0 {
		kk := (*arr)[l/2] + (*arr)[l/2-1]
		if kk%2 == 0 {
			return fmt.Sprintf("%d", kk/2)
		} else {
			return fmt.Sprintf("%d", kk/2) + ".5"
		}
	} else {
		return fmt.Sprintf("%d", (*arr)[l/2])
	}
}

func main() {
	//fmt.Println(mid(&[]int{2, 3, 5}))
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		n := 0
		_, _ = fmt.Scanf("%d", &n)
		a, b := 0, 0
		aa, bb := make([]int, 0), make([]int, 0)
		amap := make(map[int]int)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scanf("%d", &a)
			aa = append(aa, a)
			amap[j] = a
		}
		for j := 0; j < n-1; j++ {
			_, _ = fmt.Scanf("%d", &b)
			bb = append(bb, b)
		}
		pp := 0
		for {
			fmt.Printf("%s ", mid(&aa))
			aa = popA(&aa, amap[bb[pp]-1])
			pp++
			if pp == len(bb) {
				break
			}
		}
		fmt.Printf("%d", aa[0])
	}
}

//2
//5
//2 2 1 3 5
//3 1 2 5
//4
//1 1 1 1
//1 2 3
