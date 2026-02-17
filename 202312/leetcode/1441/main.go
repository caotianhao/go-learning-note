package main

import "fmt"

func buildArray1441(target []int, n int) []string {
	stackString := make([]string, 0)
	if n > target[len(target)-1] {
		n = target[len(target)-1]
	}
	for i := 0; i < n; i++ {
		if isIn1441(target, i+1) {
			stackString = append(stackString, "Push")
		} else {
			stackString = append(stackString, "Push", "Pop")
		}
	}
	return stackString
}

func isIn1441(str []int, n int) bool {
	for i := range str {
		if str[i] == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(buildArray1441([]int{1, 3}, 3))
	fmt.Println(buildArray1441([]int{1, 2, 3}, 3))
	fmt.Println(buildArray1441([]int{1, 2}, 4))
	fmt.Println(buildArray1441([]int{1, 2, 3, 4}, 4))
}
