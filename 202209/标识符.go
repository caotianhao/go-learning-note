package main

import (
	"fmt"
	"gogogo/project02"
)

// 变量名，常量名，函数名首字母大写，则表示public
func main() {
	fmt.Println(project02.HeroName)
	fmt.Println(project02.Add(12, 34))
	fmt.Println(project02.Plus(2, 3))
	project02.PrintHello()
	fmt.Println(7 << 3)
	a := 10
	b := 20
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("a =", a, "b =", b)
}
