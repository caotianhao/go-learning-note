package main

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) (res []int) {
	// 辅助数组存储一个 [2]int，对应下标和数值
	var tmp [][2]int
	for i, v := range nums {
		tmp = append(tmp, [2]int{i, v})
	}
	// 按数值降序排序
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1]
	})
	// 取前 k 个按照下标升序排序
	sort.Slice(tmp[:k], func(i, j int) bool {
		return tmp[i][0] < tmp[j][0]
	})
	// 输出
	for i := 0; i < k; i++ {
		res = append(res, tmp[i][1])
	}
	return res
}

func main() {
	fmt.Println(maxSubsequence([]int{50, -75}, 2))
	fmt.Println(maxSubsequence([]int{-1, -2, 4, 3}, 3))
	fmt.Println(maxSubsequence([]int{2, 1, 3, 3}, 2))
	fmt.Println(maxSubsequence([]int{3, 4, 3, 3, 3, 3, 3}, 2))
}
