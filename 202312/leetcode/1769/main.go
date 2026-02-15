package main

import "fmt"

func minOperations1769(boxes string) []int {
	l := len(boxes)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if boxes[j] == '1' {
				ret[i] += abs1769(i, j)
			}
		}
	}
	return ret
}

func abs1769(i, j int) int {
	if i-j > 0 {
		return i - j
	}
	return j - i
}

func main() {
	fmt.Println(minOperations1769("110"))
	fmt.Println(minOperations1769("001011"))
	fmt.Println(minOperations1769("0"))
	//s := "110"
	//fmt.Println(s[0])
}
