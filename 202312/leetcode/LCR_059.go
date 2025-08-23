package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type KthLargest2059 struct {
	sort.IntSlice
	k int
}

func Constructor2059(k int, nums []int) KthLargest2059 {
	res := KthLargest2059{k: k}
	for _, v := range nums {
		res.Add(v)
	}
	return res
}

func (kl *KthLargest2059) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func (kl *KthLargest2059) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest2059) Pop() interface{} {
	res := kl.IntSlice[len(kl.IntSlice)-1]
	kl.IntSlice = kl.IntSlice[:len(kl.IntSlice)-1]
	return res
}

func main() {
	c := Constructor2059(12, []int{1, 2, 3})
	fmt.Println(c.Add(1))
}
