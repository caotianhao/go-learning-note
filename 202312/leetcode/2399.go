package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	l, yes := len(s), 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] && j-i-1 == distance[s[i]-97] {
				yes++
			}
		}
	}
	if yes == l/2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(checkDistances("abaccb",
		[]int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa",
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
