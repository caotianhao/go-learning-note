package main

import (
	"errors"
	"fmt"
)

func div1(a, b int) int {
	return a / b
}

func errorHandle() {
	//默认情况下，程序遇到 panic 就会崩溃
	//而我们希望可以继续执行下去，还可以在遇到错误时通过邮件等方式通知管理员等
	//这就引出了错误处理机制
	//Go 中使用 panic, defer, recover 处理
	//其中，recover 是内置函数，可以捕获到错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err =", err)
			fmt.Println("除 0 错误反馈给管理员...")
		}
	}()
	i, j := 1, 0
	k := div1(i, j)
	fmt.Println("k =", k)
}

// 自定义错误
func judgeFileName(name string) error {
	if name == "cth.ini" {
		return nil
	} else {
		myErr := errors.New("big fatal err that you aren't cth")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err =", err)
				fmt.Println("文件错误反馈给管理员...")
			}
		}()
		panic(myErr)
	}
}

func main() {
	errorHandle()
	fmt.Println("---------调用 errorHandle() 之后的语句---------")
	//filename := "cth.jpg"
	filename2 := "cth.exe"
	//filename3 := "cth.ini"
	//filename4 := "cth.dst"
	//fmt.Println(judgeFileName(filename))
	fmt.Println(judgeFileName(filename2))
	//fmt.Println(judgeFileName(filename3))
	//fmt.Println(judgeFileName(filename4))
}
