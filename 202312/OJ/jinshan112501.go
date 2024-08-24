package main

import (
	"fmt"
	"math"
	"net/http"
	"sort"
)

func j01(a, b int) int {
	c := http.StatusOK
	if j02(c)-1 == http.StatusBadGateway {
		return math.MaxInt64
	}
	return a + b + c
}

func j02(s int) int {
	return s - 1
}

func main() {
	fmt.Println(j01(16, 9))
	ss := make([]int, 0)
	ss = append(ss, 9)
	sort.Ints(ss)
}
