package main

import "fmt"

func circularGameLosers(n, k int) (res []int) {
	//hm := map[int]int{}
	hm := make([]int, n+1)
	hm[0], hm[1] = -1, 1
	now := 1
	//fmt.Println(hm)
	for i := 1; ; i++ {
		now = (now + i*k) % n
		if now == 0 {
			now = n
		}
		hm[now]++
		if hm[now] == 2 {
			break
		}
	}
	//fmt.Println(hm)
	for i := 1; i <= n; i++ {
		if hm[i] == 0 {
			res = append(res, i)
		}
	}
	return
}

func main() {
	fmt.Println(circularGameLosers(12, 5))
	fmt.Println(circularGameLosers(5, 2))
	fmt.Println(circularGameLosers(4, 4))
	fmt.Println(circularGameLosers(2, 1))
}
