package main

import (
	"container/heap"
	"fmt"
)

//type myHeap struct {
//	sort.IntSlice
//}

type myHeap []int64

//func (h *myHeap) Push(i interface{}) {
//	h.IntSlice = append(h.IntSlice, i.(int))
//}
//
//func (h *myHeap) Pop() interface{} {
//	o := h.IntSlice[len(h.IntSlice)-1]
//	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
//	return o
//}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(i interface{}) {
	*h = append(*h, i.(int64))
}

func (h *myHeap) Pop() interface{} {
	o := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return o
}

/*
5
6 3
8 1 10 1 1 1
5 3
1 13 5 12 3
10 6
10 8 18 13 3 8 6 4 14 12
10 5
9 9 2 11 14 33 4 9 14 12
1 1
1
*/

/*
5
6 3
8 1 10 1 1 1
5 3
1 13 5 12 3
10 6
10 8 18 13 3 8 6 4 14 12
10 4
9 9 2 11 14 33 4 9 14 12
1 1
1
*/

// 通过60%
// 问题在于
// 10 4
// 9 9 2 11 14 33 4 9 14 12
// 这样的用例
// 应该数组为 9 2 4 9
// 但是在实现上暂时想不到办法去剔除哪一个9
// 导致实际数组变为了 9 9 2 4
// 进而导致答案错误
func main() {
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	n, k := 0, 0
	var ti int64
	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%d %d", &n, &k)
		needTime := make([]int64, 0)
		hh := &myHeap{}
		heap.Init(hh)
		m := make(map[int64]int)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scanf("%d", &ti)
			needTime = append(needTime, ti)
			heap.Push(hh, ti)
		}
		for j := 0; j < k; j++ {
			m[heap.Pop(hh).(int64)]++
		}
		trueTaskList := make([]int64, 0)
		for j := 0; j < len(needTime); j++ {
			if v, ok := m[needTime[j]]; ok && v > 0 {
				m[needTime[j]]--
				trueTaskList = append(trueTaskList, needTime[j])
			}
		}
		//fmt.Println("ttl", trueTaskList)
		//if len(trueTaskList) != k {
		//	diff := len(trueTaskList) - k
		//	trueTaskList = trueTaskList[diff:]
		//}
		trLen := len(trueTaskList)
		if trLen == 1 {
			fmt.Println(trueTaskList[0])
		} else if trLen == 2 {
			fmt.Println(max(trueTaskList[0], trueTaskList[1]))
		} else {
			var s int64
			for _, v := range trueTaskList {
				s += v
			}
			h2 := &myHeap{}
			heap.Init(h2)
			prefix := make([]int64, trLen+1)
			for p := 0; p < trLen; p++ {
				prefix[p+1] += prefix[p] + trueTaskList[p]
			}
			//fmt.Println("pre", prefix)
			for p := 1; p < trLen; p++ {
				heap.Push(h2, max(s-prefix[p], prefix[p]))
			}
			fmt.Println(heap.Pop(h2).(int64))
		}
	}
}
