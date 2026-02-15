package main

import "fmt"

func singleNumber137(nums []int) int {
	// 一个元素出现一次，其他三次
	// 思路：依次确定答案的每一位
	// 由于其他数出现三次，那么数组全部数字每一位之和必为 3 或 4
	// 计算对 3 的余数即可知道答案的对应那一位为 0 或 1
	// 这里必须要用 32 位，否则是默认 64 位无法输出负数
	var ans int32
	for i := 0; i < 32; i++ {
		var cnt int32
		for j := 0; j < len(nums); j++ {
			cnt += 1 & (int32(nums[j]) >> i)
		}
		if cnt%3 == 1 {
			ans += 1 << i
			// 或 ans |= 1 << i
		}
	}
	return int(ans)
}

func main() {
	fmt.Println(singleNumber137([]int{0, 1, 0, 1, 0, 1, 100, 102, 102, 100, 102, 100, 897, 6, 6, 3, 3, 3, 6}))
	fmt.Println(singleNumber137([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
