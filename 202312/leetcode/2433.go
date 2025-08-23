package main

import "fmt"

func findArray(pref []int) []int {
	l := len(pref)
	res := make([]int, l)
	res[0] = pref[0]
	for i := 1; i < l; i++ {
		res[i] = pref[i] ^ pref[i-1]
	}
	return res
}

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
	fmt.Println(findArray([]int{12}))
}
