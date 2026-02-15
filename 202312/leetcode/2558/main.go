package main

import (
	"container/heap"
	"fmt"
	"math"
)

type heap2558 []int

func (h *heap2558) Len() int {
	return len(*h)
}

func (h *heap2558) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap2558) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *heap2558) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *heap2558) Pop() any {
	o := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return o
}

func pickGifts1(gifts []int, k int) int64 {
	ans := 0
	for i := 0; i < k; i++ {
		maxGift, ind := -1, -1
		for k, v := range gifts {
			if v > maxGift {
				maxGift = v
				ind = k
			}
		}
		gifts[ind] = int(math.Sqrt(float64(gifts[ind])))
	}
	for _, v := range gifts {
		ans += v
	}
	return int64(ans)
}

func pickGifts(gifts []int, k int) (ans int64) {
	h := heap2558(gifts)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		s := heap.Pop(&h).(int)
		heap.Push(&h, int(math.Sqrt(float64(s))))
	}
	for h.Len() != 0 {
		ans += int64(heap.Pop(&h).(int))
	}
	return
}

func main() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4))
	fmt.Println(pickGifts1([]int{25, 64, 9, 4, 100}, 4))
}
