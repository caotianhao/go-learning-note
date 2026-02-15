package main

import "fmt"

//每一个源文件都可以包含一个init函数
//该函数在main()之前调用，用于初始化变量等工作
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
