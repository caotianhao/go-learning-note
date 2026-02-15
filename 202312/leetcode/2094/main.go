package main

import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {
	m := map[int]struct{}{}
	ans, l := make([]int, 0), len(digits)
	for i := 0; i < l; i++ {
		if digits[i]%2 == 0 {
			for j := 0; j < l; j++ {
				if j == i {
					continue
				}
				for k := 0; k < l; k++ {
					if k == i || k == j {
						continue
					}
					if digits[j] != 0 && digits[k] != 0 {
						m[digits[j]*100+digits[k]*10+digits[i]] = struct{}{}
						m[digits[k]*100+digits[j]*10+digits[i]] = struct{}{}
					} else if digits[j] == 0 && digits[k] != 0 {
						m[digits[k]*100+digits[i]] = struct{}{}
					} else if digits[j] != 0 && digits[k] == 0 {
						m[digits[j]*100+digits[i]] = struct{}{}
					}
				}
			}
		}
	}
	for i := range m {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}

func main() {
	//fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
}
