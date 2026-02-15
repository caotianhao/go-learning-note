package main

import (
	"fmt"
)

func main() {
	//switch后面可以不带东西，当if-else用
	var i int
	fmt.Print("input i: ")
	_, _ = fmt.Scanln(&i)
	switch {
	case i > 10:
		fmt.Println("i > 10")
	case i <= 10:
		fmt.Println("i <= 10")
	}

	//switch后面可以直接声明或定义一个变量，分号结束，不推荐
	switch a := 1; {
	case i+a != 25:
		fmt.Println("i+a != 25")
	case i+a == 25:
		fmt.Println("i+a == 25")
	}

	//switch有fallthrough，会继续执行一个case
	var j int
	fmt.Print("input j: ")
	_, _ = fmt.Scanln(&j)
	switch {
	case j%2 == 0:
		fmt.Println("j是偶数")
		fallthrough
	case j%3 == 0:
		fmt.Println("j是3的倍数")
	case j%5 == 0:
		fmt.Println("j是5的倍数")
	default:
		fmt.Println("No")
	}

}
