package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maximumValue(strs []string) int {
	sli := make([]int, 0)
	for _, v := range strs {
		if isChar2496(v) {
			sli = append(sli, len(v))
		} else {
			tmp, _ := strconv.ParseInt(v, 10, 64)
			sli = append(sli, int(tmp))
		}
	}
	sort.Ints(sli)
	return sli[len(sli)-1]
}

func isChar2496(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(maximumValue([]string{"alic3", "bob", "3", "4", "00000"}))
}
