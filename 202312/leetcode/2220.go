package main

import "fmt"

func minBitFlips(start int, goal int) int {
	startSlice := decToBin(start)
	goalSlice := decToBin(goal)
	ls, lg, cnt := len(startSlice), len(goalSlice), 0
	if ls > lg {
		for i := 0; i < ls-lg; i++ {
			goalSlice = append(goalSlice, 0)
		}
		lg = ls
	} else if lg > ls {
		for i := 0; i < lg-ls; i++ {
			startSlice = append(startSlice, 0)
		}
		ls = lg
	}
	for i := 0; i < ls; i++ {
		if startSlice[i] != goalSlice[i] {
			cnt++
		}
	}
	return cnt
}

func decToBin(n int) []int {
	if n == 0 {
		return []int{0}
	}
	bin := make([]int, 0)
	for ; n != 1; n /= 2 {
		bin = append(bin, n%2)
	}
	bin = append(bin, 1)
	return bin
}

func main() {
	fmt.Println(minBitFlips(3, 4))
	fmt.Println(decToBin(0))
}
