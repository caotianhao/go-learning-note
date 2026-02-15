package main

import (
	"fmt"
	"sync"
	"time"
)

var goroutineSumMap = make(map[int]int)

// 全局互斥锁
var lock sync.Mutex

func goroutineSum(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	//加锁，防止一起写
	lock.Lock()
	goroutineSumMap[n] = res
	//加锁后要解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 2000; i++ {
		//这里若是直接 goroutineSum(i) 就是正常的做法，但我们想通过 goroutine 加速
		//goroutineSum(i)
		//但是这样直接 go goroutineSum(i) 却得不到想要的结果
		//其实是相当于很多个协程在同时执行写操作，需要一个全局互斥锁
		go goroutineSum(i)
	}
	//这个得加，不然在这个时间内没执行完，所以会出现不全的情况！
	//只要时间能保证足够运行完就行，不然主线程结束协程没完事！
	//注意！！！
	time.Sleep(time.Second * 2)
	//由于 map 无序，所以输出并不是从小到大排的

	//这对锁不加的话还是会引起底层资源竞争
	//在遍历 map 时我们已经知道我们设定的时间可以完成协程
	//但底层和主线程不知道，所以得加
	//可以用 race 看资源竞争是否存在
	lock.Lock()
	for i := 1; i < len(goroutineSumMap); i++ {
		fmt.Println(i, "->", goroutineSumMap[i])
	}
	lock.Unlock()
}
