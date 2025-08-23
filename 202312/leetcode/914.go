package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 1 {
		return false
	}
	cardMap, cnt := map[int]int{}, make([]int, 0)
	for _, v := range deck {
		cardMap[v]++
	}
	for _, v := range cardMap {
		cnt = append(cnt, v)
	}
	if len(cnt) == 1 {
		return true
	}
	for i := 0; i < len(cnt); i++ {
		for j := 0; j < len(cnt); j++ {
			if gcd914(cnt[i], cnt[j]) == 1 {
				return false
			}
		}
	}
	return true
}

func gcd914(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	//fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3}))
	//fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 1, 1, 1}))
	//fmt.Println(hasGroupsSizeX([]int{0}))
	fmt.Println(hasGroupsSizeX([]int{2, 2}))
	fmt.Println(hasGroupsSizeX([]int{2, 3}))
	fmt.Println(gcd914(18, 12), gcd914(12, 18))
}
