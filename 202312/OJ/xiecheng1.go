package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type myHee struct {
	sort.IntSlice
}

func (h *myHee) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *myHee) Pop() interface{} {
	o := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return o
}

func nextA(nums []int) (r []int) {
	l := len(nums)
	hh := &myHee{}
	heap.Init(hh)
	for i := 1; i <= l; i++ {
		heap.Push(hh, i)
	}
	for i := 0; i < l; i++ {
		r = append(r, heap.Pop(hh).(int))
	}
	return
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	arr := make([]int, 0)
	a := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	q := nextA(arr)
	for _, v := range q {
		fmt.Printf("%d ", v)
	}
}
