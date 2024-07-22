package main

import (
	"fmt"
)

func unCurry(a, b, c int) int {
	return a + b + c
}

// 柯里化
// 是一种将接受多个参数的函数转换成一系列接受单一参数的函数的技术
// 它的目标是让函数的使用更加灵活，可以通过部分应用参数来生成新的函数
func curry(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func main() {
	fmt.Println(unCurry(2, 3, 4)) // 9
	fmt.Println(curry(2)(3)(4))   // 9
}
