package main

import "fmt"

//全局匿名函数，视频中是大写Func1
var func1 = func(n1, n2 int) int {
	return n1 * n2
}

func main() {
	//在定义匿名函数的时候就调用，这样只能用一次
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1 =", res1)

	//将匿名函数赋给变量，可以多次调用
	a := func(n1, n2 int) int {
		return n1 - n2
	}
	res2 := a(15, 5)
	fmt.Println("res2 =", res2)
	res3 := a(16, 5)
	fmt.Println("res3 =", res3)

	//调用全局匿名函数
	res4 := func1(20, 30)
	fmt.Println("res4 =", res4)
}
