package main

import "fmt"

func main() {
	var x interface{}
	y := 1.23
	//把float64赋给空接口，因为空接口可以接收任意类型的值
	fmt.Printf("x type is %T, x is %v\n", x, x)
	x = y
	fmt.Printf("x type is %T, x is %v\n", x, x)
	if a, ok := x.(float64); ok {
		fmt.Printf("a type is %T, a is %v\n", a, a)
	}
}
