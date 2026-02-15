package main

import (
	"fmt"
	"time"
)

func printHello() {
	for i := 1; i <= 10; i++ {
		fmt.Println("out " + fmt.Sprintf("%d", i))
		//1s
		time.Sleep(time.Second)
	}
}

func main() {
	//开启一个协程
	go printHello()
	//主线程结束了，协程也结束
	//协程结束了，主线程没完成的话可以继续
	for i := 1; i <= 10; i++ {
		fmt.Println("main " + fmt.Sprintf("%d", i))
		//0.2s
		time.Sleep(time.Second / 5)
	}
}
