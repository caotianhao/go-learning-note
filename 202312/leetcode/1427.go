package main

import "fmt"

func stringShift(s string, shift [][]int) string {
	for _, v := range shift {
		if v[0] == 0 {
			for i := 0; i < v[1]; i++ {
				s = shiftLeft(s)
			}
		} else {
			for i := 0; i < v[1]; i++ {
				s = shiftRight(s)
			}
		}
	}
	return s
}

func shiftLeft(s string) string {
	return s[1:] + string(s[0])
}

func shiftRight(s string) string {
	return string(s[len(s)-1]) + s[:len(s)-1]
}

func main() {
	fmt.Println(stringShift("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}))
}
