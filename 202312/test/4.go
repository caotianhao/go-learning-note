package main

import (
	"fmt"
)

func main() {
	n := 5 // 选取的数字范围为1到n
	s := 9 // 目标和为s
	combinations := make([][]int, 0)
	currentCombination := make([]int, 0)
	findCombinations(n, s, 1, currentCombination, &combinations)
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
func findCombinations(n, s, start int, currentCombination []int, combinations *[][]int) {
	if s == 0 {
		// 找到一个满足条件的组合
		*combinations = append(*combinations, append([]int{}, currentCombination...))
		return
	}
	for i := start; i <= n; i++ {
		if i <= s {
			// 将当前数字添加到组合中
			currentCombination = append(currentCombination, i)
			// 递归查找下一个数字
			findCombinations(n, s-i, i+1, currentCombination, combinations)
			// 移除最后一个数字，继续搜索其他组合
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
}
