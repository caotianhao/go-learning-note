package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//返回从1970开始到现在的纳秒数
	rand.NewSource(time.Now().UnixNano())
	time0 := time.Now()
	cnt := 0
	for {
		//[0,10000)+1
		a := rand.Intn(10000) + 1
		cnt++
		//fmt.Println(a)
		if a == 1234 {
			break
		}
	}
	time1 := time.Since(time0)
	fmt.Println(time1)
	fmt.Println(cnt)
}
