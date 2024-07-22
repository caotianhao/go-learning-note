package main

import (
	"fmt"
	"strings"
)

//判断传入的文件名是否含有后缀，如果有就直接返回，没有就加上
func judgeSuffix(suffix string) func(string) string {
	//返回的匿名函数与 suffix 构成一个整体，称为闭包
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//好处在于，使用闭包的时候，每次只需开始前传入一次后缀，之后每次调用只需要传一个参数
	//传统方法需要每次传入两个参数，终于理解了
	f := judgeSuffix(".jpg")
	fmt.Println(f("winter"))   //winter.jpg
	fmt.Println(f("bird.jpg")) //bird.jpg
	fmt.Println(f("cat.mp3"))  //cat.mp3.jpg
	fmt.Println(f("cat.jpp"))  //cat.jpp.jpg
}
