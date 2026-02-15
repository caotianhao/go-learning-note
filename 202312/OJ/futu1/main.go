package main

import "fmt"

func f(x int64, b *[]int64) int {
	// 问题转化为要求 (b_1 - x)(b_2 - x)...(b_m - x) > 0
	// cnt 计算在乘积中的每项小于 0 的个数
	fu := 0
	for i := 0; i < len(*b); i++ {
		//fmt.Println("a", (*b)[i], x)
		// 如果在乘积当中有一项为 0，直接返回一个特殊值
		if (*b)[i]-x == 0 {
			return -1
		}
		if (*b)[i]-x < 0 {
			fu++
		}
		//if (*b)[i]-x > 0 {
		//	z++
		//}
	}
	return fu
}

func main() {
	var n, m int
	var a, b int64

	coin := make([]int64, 0)
	bb := make([]int64, 0)

	_, _ = fmt.Scanf("%d %d", &n, &m)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		coin = append(coin, a)
	}

	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d", &b)
		bb = append(bb, b)
	}
	//fmt.Println("coin", coin, "bb", bb)

	cnt := 0
	for i := 0; i < len(coin); i++ {
		t := f(coin[i], &bb)
		//fmt.Println("aaa", f(coin[i], &bb))
		if t == -1 {
			fmt.Printf("%d", 0)
			return
		}
		if t%2 == 0 {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)
}
