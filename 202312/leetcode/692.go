package main

import (
	"container/heap"
	"fmt"
)

type pair692 struct {
	word string
	cnt  int
}

type priorityQueue692 []pair692

func (h *priorityQueue692) Len() int {
	return len(*h)
}

func (h *priorityQueue692) Less(i, j int) bool {
	return (*h)[i].cnt < (*h)[j].cnt ||
		(*h)[i].cnt == (*h)[j].cnt && (*h)[i].word > (*h)[j].word
}

func (h *priorityQueue692) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *priorityQueue692) Push(x any) {
	*h = append(*h, x.(pair692))
}

func (h *priorityQueue692) Pop() any {
	o := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return o
}

func topKFrequent1(words []string, k int) []string {
	hm := make(map[string]int)
	for _, v := range words {
		hm[v]++
	}
	ans := make([]string, k)
	h := &priorityQueue692{}
	heap.Init(h)
	for str, cnt := range hm {
		heap.Push(h, pair692{str, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair692).word
	}
	return ans
}

func main() {
	fmt.Println(topKFrequent1([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent1([]string{"the", "day", "is", "sunny",
		"the", "the", "the", "sunny", "is", "is", "is"}, 4))
}
