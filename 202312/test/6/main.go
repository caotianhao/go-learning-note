package main

import "fmt"

func minOperation(nums []int) int {
	first := nums[0]
	maxC := -1
	for _, v := range nums {
		if v > maxC {
			maxC = v
		}
	} // 我想的太简单了，先试试看看能过多少...
	cnt := 0 // 技术3，安全4
	if first >= maxC {
		return 0
	} else {
		for first < maxC {
			first <<= 1
			cnt++
		}
	}
	return cnt
}

func main() {
	n, a := 0, 0
	_, _ = fmt.Scanf("%d", &n)
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	fmt.Printf("%d", minOperation(arr))
}
