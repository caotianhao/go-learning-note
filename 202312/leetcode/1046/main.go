package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

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

func lastStoneWeight(stones []int) int {
	l := len(stones)
	if l == 1 {
		return stones[0]
	}
	if l == 2 {
		if stones[0] == stones[1] {
			return 0
		} else {
			return int(math.Abs(float64(stones[0]) - float64(stones[1])))
		}
	}
	sort.Ints(stones)
	stonesSlice := stones[:l-2]
	if stones[l-1] == stones[l-2] {
		return lastStoneWeight(stonesSlice)
	} else {
		stonesSlice = append(stonesSlice, stones[l-1]-stones[l-2])
		return lastStoneWeight(stonesSlice)
	}
}

func lastStoneWeight1(stones []int) int {
	h := &maxHeap{}
	heap.Init(h)
	for _, v := range stones {
		heap.Push(h, v)
	}
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		if a > b {
			heap.Push(h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	} else {
		return heap.Pop(h).(int)
	}
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight1([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1, 16, 54, 41, 541, 412, 223, 861, 65, 20, 14, 1, 3}))
}
