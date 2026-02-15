package main

import (
	"fmt"
	"strconv"
)

func summaryRanges1(nums []int) (res []string) {
	l, p, q := len(nums), 0, 1
	if l == 0 {
		return
	}
	if l == 1 {
		return []string{strconv.FormatInt(int64(nums[0]), 10)}
	}
	for q < l && p < l {
		if nums[q]-nums[p] == q-p {
			q++
		} else {
			if p == q-1 {
				res = append(res, strconv.FormatInt(int64(nums[p]), 10))
			} else {
				res = append(res, strconv.FormatInt(int64(nums[p]), 10)+
					"->"+strconv.FormatInt(int64(nums[q-1]), 10))
			}
			p = q
			q++
		}
	}
	if p == q-1 {
		res = append(res, strconv.FormatInt(int64(nums[p]), 10))
	} else {
		res = append(res, strconv.FormatInt(int64(nums[p]), 10)+
			"->"+strconv.FormatInt(int64(nums[q-1]), 10))
	}
	return
}

func summaryRanges(nums []int) (res []string) {
	l := len(nums)
	if l == 0 {
		return
	}
	p, q := 0, 1
	for q < l {
		if nums[q]-nums[p] == q-p {
			q++
		} else {
			if p == q-1 {
				res = append(res, fmt.Sprintf("%d", nums[p]))
			} else {
				res = append(res, fmt.Sprintf("%d", nums[p])+"->"+fmt.Sprintf("%d", nums[q-1]))
			}
			p = q
			q++
		}
	}
	if p == q-1 {
		res = append(res, fmt.Sprintf("%d", nums[p]))
	} else {
		res = append(res, fmt.Sprintf("%d", nums[p])+"->"+fmt.Sprintf("%d", nums[q-1]))
	}
	return
}

func main() {
	fmt.Println(summaryRanges1([]int{0, 1, 2, 4, 5, 7}))    // "0->2","4->5","7"
	fmt.Println(summaryRanges1([]int{0, 2, 3, 4, 6, 8, 9})) // "0","2->4","6","8->9"
	fmt.Println(summaryRanges([]int{-2, -1, 0, 1, 3, 4}))
	fmt.Println(summaryRanges([]int{-2, 544}))
	fmt.Println(summaryRanges([]int{-2}))
	fmt.Println(summaryRanges([]int{}))
}
