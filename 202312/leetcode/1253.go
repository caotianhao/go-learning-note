package main

import "fmt"

func reconstructMatrix(upper, lower int, s []int) [][]int {
	l := len(s)
	up, low := make([]int, l), make([]int, l)
	vu, vl := upper, lower
	for i, v := range s {
		if v == 2 {
			up[i], low[i] = 1, 1
			upper--
			lower--
		}
	}
	for i, v := range s {
		if v == 1 {
			if upper > 0 {
				up[i] = 1
				upper--
			} else if lower > 0 {
				low[i] = 1
				lower--
			}
		}
	}
	verify := func(vu, vl int, up, low, s []int) bool {
		for _, v := range up {
			if v == 1 {
				vu--
			}
		}
		if vu != 0 {
			return false
		}
		for _, v := range low {
			if v == 1 {
				vl--
			}
		}
		if vl != 0 {
			return false
		}
		for i, v := range s {
			if up[i]+low[i] != v {
				return false
			}
		}
		return true
	}
	if verify(vu, vl, up, low, s) {
		return [][]int{up, low}
	}
	return [][]int{}
}

func main() {
	fmt.Println(reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}))
	fmt.Println(reconstructMatrix(2, 3, []int{2, 2, 1, 1}))
}
