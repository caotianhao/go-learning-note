package main

import (
	"container/heap"
	"fmt"
	"math"
)

type h2530 []int64

func (h *h2530) Len() int {
	return len(*h)
}

func (h *h2530) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *h2530) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *h2530) Push(i any) {
	*h = append(*h, i.(int64))
}

func (h *h2530) Pop() any {
	o := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return o
}

func maxKelements(nums []int, k int) (score int64) {
	h := &h2530{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, int64(num))
	}
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(int64)
		score += p
		heap.Push(h, int64(math.Ceil(float64(p)/3.0)))
	}
	return score
}

func main() {
	fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5))
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}
