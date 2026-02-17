package main

import "fmt"

func findThePrefixCommonArray(A, B []int) []int {
	res, myMap := make([]int, len(A)), map[int]int{}
	for i := range A {
		myMap[A[i]]++
		myMap[B[i]]++
		cnt := 0
		for _, v := range myMap {
			if v == 2 {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
}
