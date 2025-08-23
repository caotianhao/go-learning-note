package main

import (
	"fmt"
	"sort"
)

// 错误做法是因为把奶酪想成了 2n 块
// 其实只有 n 块奶酪，不要被两个 n 长度的数组迷惑
//func miceAndCheese(reward1 []int, reward2 []int, k int) int {
//	res, l := 0, len(reward1)
//	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
//	sort.Sort(sort.Reverse(sort.IntSlice(reward2)))
//	for i := 0; i < k; i++ {
//		res += reward1[i]
//	}
//	for i := 0; i < l-k; i++ {
//		res += reward2[i]
//	}
//	return res
//}

// 正确做法
func miceAndCheese(reward1, reward2 []int, k int) int {
	res, l := 0, len(reward1)
	diff := make([]int, l)
	for i := 0; i < l; i++ {
		// 先假设 n 块奶酪都被第 2 只老鼠吃掉
		res += reward2[i]
		// 然后建立一个差数组，表示每一块的奶酪被 1 吃掉和被 2 吃掉对分数的影响
		diff[i] = reward1[i] - reward2[i]
	}
	// 排序，选 k 个即可
	sort.Sort(sort.Reverse(sort.IntSlice(diff)))
	for i := 0; i < k; i++ {
		res += diff[i]
	}
	return res
}

func main() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))       //15
	fmt.Println(miceAndCheese([]int{1, 1}, []int{1, 1}, 2))                   //2
	fmt.Println(miceAndCheese([]int{1, 4, 4, 6, 4}, []int{6, 5, 3, 6, 1}, 1)) //24
}
