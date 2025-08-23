package main

import (
	"fmt"
	"strconv"
)

func lastVisitedIntegers(words []string) (result []int) {
	nums, rev := make([]string, 0), make([]string, 0)
	tsu := 0
	for _, v := range words {
		if v != "prev" {
			tsu = 0
			nums = append(nums, v)
			rev = reverse2899(nums)
		} else {
			tsu++
			if tsu-1 < len(rev) {
				tmp, _ := strconv.Atoi(rev[tsu-1])
				result = append(result, tmp)
			} else {
				result = append(result, -1)
			}
		}
	}
	return
}

func reverse2899(arr []string) (r []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		r = append(r, arr[i])
	}
	return
}

func main() {
	//fmt.Println(lastVisitedIntegers([]string{"1", "2", "prev", "prev", "prev"}))
	//fmt.Println(lastVisitedIntegers([]string{"1"}))
	//fmt.Println(lastVisitedIntegers([]string{"prev"}))
	fmt.Println(lastVisitedIntegers([]string{"prev", "prev", "94", "56", "prev", "32", "prev", "prev", "prev"}))
	fmt.Println(reverse2899([]string{"1", "2", "3"}))
}
