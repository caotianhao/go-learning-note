package main

import (
	"fmt"
	"sort"
)

func abs1(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	var n, num, sum, a, b int
	_, _ = fmt.Scanf("%d", &n)
	arr := make([]int, 0)
	hashMap := make(map[int]int)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		arr = append(arr, num)
		hashMap[num]++
	}
	for _, v := range arr {
		sum += v
	}
	if k := sum % n; k == 0 {
		flag := sum / n
		for _, v := range arr {
			if v < flag {
				a += flag - v
			}
			if v > flag {
				b += v - flag
			}
		}
		if a > b {
			fmt.Printf("%d", a)
		} else {
			fmt.Printf("%d", b)
		}
	} else {
		if len(hashMap) == 2 {
			sort.Ints(arr)
			f, l := arr[0], arr[n-1]
			if hashMap[f] > hashMap[l] && hashMap[f] > 1 && hashMap[l] > 1 {
				fmt.Printf("%d", abs1(f-l)*(hashMap[l]-1))
			}
			if hashMap[f] < hashMap[l] && hashMap[f] > 1 && hashMap[l] > 1 {
				fmt.Printf("%d", abs1(f-l)*(hashMap[f]-1))
			}
		} else {
			fmt.Printf("%d", 0)
		}
	}
}
