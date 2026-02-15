package main

import "fmt"

//按字典序输出 1~n 的全排列，例如 n=3 时：123,132,213,231,312,321
//n 全局变量
var n = 5

//p 为当前排列，11 是因为 10 的全排列基本就够慢了，正常设为 n+1 就行
var p [11]int

//hashTable 记录整数 x 是否在 p[] 中
var hashTable [11]bool

//处理排列的第 index 位
func fullPermutation(index int) {
	//递归边界，已经处理完排列的 1~n 位
	if index == n+1 {
		for i := 1; i <= n; i++ {
			fmt.Print(p[i], " ")
		}
		fmt.Println()
		return
	}
	//枚举 1~n，试图将 x 填入 p[]
	for x := 1; x <= n; x++ {
		if !hashTable[x] {
			//如果 x 不在 p[] 中，将 p 的第 index 位置为 x，即 x 加入排列
			p[index] = x
			//hashTable[x] 置为 true
			hashTable[x] = true
			//处理第 index+1 位
			fullPermutation(index + 1)
			//已处理完 p[index] 为 x 的子问题，恢复状态为 false
			hashTable[x] = false
		}
	}
}
func main() {
	fullPermutation(1)
}
