package main

import (
	"fmt"
	"sort"
)

func maxOddCount(arr []int, k int) int {
	odd := 0 // 统计原数组奇数个数
	bits := make([]int, 0)
	for _, v := range arr {
		if v&1 == 0 { // 是偶数的话
			cnt := 0
			for v&1 == 0 { // 查末尾几个0
				cnt++
				v >>= 1
			}
			bits = append(bits, cnt)
		} else {
			odd++
		}
	} // 此时 对应 114514 的 bits 为 002002，这里去掉0，所以 bits -> 22
	sort.Ints(bits)
	for i := 0; i < len(bits); i++ {
		if k < bits[0] {
			return odd
		}
		k -= bits[i] // 现在应该对了
		if k == 0 {
			return odd + i + 1
		}
		if k < 0 {
			return odd + i
		}
	}
	return 0
}

func main() {
	n, k := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &k)
	arr := make([]int, 0)
	num := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	fmt.Printf("%d", maxOddCount(arr, k))
}
