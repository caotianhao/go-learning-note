package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	l := len(nums)
	first, last := -1, l
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			first = i
			break
		}
	}
	for i := l - 1; i >= 0; i-- {
		if nums[i] == 1 {
			last = i
			break
		}
	}
	if first == -1 || last == l || first == last {
		return true
	}
	p, q := first, first+1
	for p != last {
		if nums[q] == 1 {
			if q-p-1 < k {
				return false
			}
			p = q
		}
		q++
	}
	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1, 1, 0}, 2))
	//fmt.Println(kLengthApart([]int{1, 1, 1, 1, 1}, 0))
	//fmt.Println(kLengthApart([]int{0, 1, 0, 1}, 1))
}
