package main

import "fmt"

// 正常来讲写成 [t int | float64 | string] 就行
// 但是当涉及指针类型的编译器就会解析为乘号
func add[t interface{ int | float64 | string }](a, b t) t {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add("hel", "lo"))

	// 由编译器自己推断类型
	fmt.Println(add(1.3, 2.6))

	// 显式指定
	fmt.Println(add[float64](1.3, 2.6))
}
