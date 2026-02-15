package main

import (
	"container/heap"
	"fmt"
	"sort"
)

const mod2 int = 1e9 + 7

type maxHeap struct {
	sort.IntSlice
}

func (h *maxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *maxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	r := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return r
}

func getMaxSum(nums []int, k int) int {
	h := &maxHeap{}
	sum := 0
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	for i := 0; i < k; i++ {
		max1 := heap.Pop(h).(int)
		max2 := heap.Pop(h).(int)
		max0 := (max1 * max2) % mod2
		heap.Push(h, max0)
		heap.Push(h, 1)
	}
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
		sum = sum % mod2
	}
	return sum % mod2
}

func main() {
	n, k := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &k)
	a := 0
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	fmt.Printf("%d", getMaxSum(arr, k))
}
