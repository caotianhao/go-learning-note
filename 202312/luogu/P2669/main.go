package main

import (
	"fmt"
	"sort"
)

const c2669 = 10001

func main() {
	k := 0
	_, _ = fmt.Scan(&k)
	ori := make([]int, 0)
	for i := 0; i < c2669; i++ {
		ori = append(ori, i+1)
	}
	prefix := make([]int, c2669)
	prefix[0] = ori[0]
	for i := 1; i < c2669; i++ {
		prefix[i] = prefix[i-1] + ori[i]
	}
	idx := sort.Search(len(prefix), func(i int) bool {
		return prefix[i] > k
	})
	diff := k - prefix[idx-1]
	//fmt.Println(idx, diff)
	coin := 0
	for i := 1; i <= idx; i++ {
		for j := 0; j < i; j++ {
			coin += i
		}
	}
	for i := 0; i < diff; i++ {
		coin += idx + 1
	}
	fmt.Println(coin)
}
