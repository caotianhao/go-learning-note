package main

import "fmt"

type MRUQueue struct {
	arr []int
}

func Constructor1756(n int) MRUQueue {
	arr := make([]int, 0)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return MRUQueue{arr: arr}
}

func (mru *MRUQueue) Fetch(k int) int {
	r := mru.arr[k-1]
	mru.arr = move1756(mru.arr, k-1)
	return r
}

func move1756(arr []int, index int) []int {
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if i == index {
			continue
		}
		newArr = append(newArr, arr[i])
	}
	newArr = append(newArr, arr[index])
	return newArr
}

func main() {
	mRUQueue := Constructor1756(8) // 初始化队列为 [1,2,3,4,5,6,7,8]。
	fmt.Println(mRUQueue.Fetch(3))
	fmt.Println(mRUQueue.Fetch(5))
	fmt.Println(mRUQueue.Fetch(2))
	fmt.Println(mRUQueue.Fetch(8)) // 3622
}
