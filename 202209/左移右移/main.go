package main

import "fmt"

func main() {
	//左移：符号位不变，低位补0
	fmt.Println(-5 << 2)
	//右移：符号位不变，用符号位补充溢出的高位
	fmt.Println(-5 >> 2)
}
