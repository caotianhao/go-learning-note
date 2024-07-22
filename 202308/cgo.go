package main

/*
#include <stdio.h>
void cgoTest() {
	printf("Today is %d year.\n", 2023);
}
*/
import "C"

// C 代码写在注释里，且要在 import 之前，且注释完要在下一行单独地 import "C"
func main() {
	// 调用
	C.cgoTest()
	// 可以使用 go tool cgo xxx.go 命令查看中间文件
}
