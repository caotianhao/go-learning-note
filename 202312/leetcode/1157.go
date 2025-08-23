package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 超时 超时 超时 超时 超时 超时

type MajorityChecker struct {
	arr []int
}

func Constructor1157(arr []int) MajorityChecker {
	return MajorityChecker{arr}
}

func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	tmpArr := mc.arr[left : right+1]
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		num := tmpArr[rand.Int()%(right+1-left)]
		cnt := 0
		for _, v := range tmpArr {
			if v == num {
				cnt++
			}
			if cnt >= threshold {
				return num
			}
		}
	}
	return -1
}

func main() {
	c := Constructor1157([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(c.Query(0, 5, 4)) //1
	fmt.Println(c.Query(0, 3, 3)) //-1
	fmt.Println(c.Query(2, 3, 2)) //2
}
