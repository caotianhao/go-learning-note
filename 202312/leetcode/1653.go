package main

import "fmt"

func minimumDeletions(s string) (ret int) {
	//1. 要么把全部的 a 删除
	//2. 要么把全部的 b 删除
	//3. 要么把左面的 b 和右面的 a 删除
	//计算任何位置左面的 b 和右面的 a
	//当位置在最左面和最右面的时候正好符合全部删除 a/b 的情况
	leftB, rightA := 0, 0
	//最左面开始，先统计 a
	for _, v := range s {
		if v == 'a' {
			rightA++
		}
	}
	//第 1 种情况
	ret = rightA
	//从左开始，第 3 种情况，最后到最后就是第 2 种情况
	for _, v := range s {
		if v == 'a' {
			rightA--
		} else {
			leftB++
		}
		//时刻记录
		ret = min(ret, rightA+leftB)
	}
	return ret
}

func main() {
	fmt.Println(minimumDeletions("abababbaabbababababababa"))
	fmt.Println(minimumDeletions("aababbab"))
}
