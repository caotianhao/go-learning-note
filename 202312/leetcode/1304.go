package main

import "fmt"

func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{1234, -1234}
	}
	what, sum := make([]int, 0), 0
	what = append(what, 0)
	for i := 1; i < n-1; i++ {
		what = append(what, i)
	}
	for i := 1; i < n-1; i++ {
		sum += i
	}
	what = append(what, -sum)
	return what
}

func main() {
	fmt.Println(sumZero(99))
}
