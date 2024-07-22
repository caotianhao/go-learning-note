package main

import (
	"container/heap"
	"fmt"
)

type PQStruct struct {
	value    int
	priority int
}

type PriorityQueue []PQStruct

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	// 按优先级从大到小排列，若优先级相同，则按照值的大小排列
	return (*pq)[i].priority > (*pq)[j].priority ||
		(*pq)[i].priority == (*pq)[j].priority && (*pq)[i].value > (*pq)[j].value
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(PQStruct))
}

func (pq *PriorityQueue) Pop() interface{} {
	o := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return o
}

func main() {
	myPriorityQueue := &PriorityQueue{}
	heap.Init(myPriorityQueue)

	heap.Push(myPriorityQueue, PQStruct{value: 1, priority: 5})
	heap.Push(myPriorityQueue, PQStruct{value: 2, priority: 3})
	heap.Push(myPriorityQueue, PQStruct{value: 3, priority: 9})
	heap.Push(myPriorityQueue, PQStruct{value: 4, priority: 1})
	heap.Push(myPriorityQueue, PQStruct{value: 5, priority: 7})
	heap.Push(myPriorityQueue, PQStruct{value: 6, priority: 7})
	heap.Push(myPriorityQueue, PQStruct{value: 7, priority: 1})
	heap.Push(myPriorityQueue, PQStruct{value: 8, priority: 11})

	for myPriorityQueue.Len() > 0 {
		item := heap.Pop(myPriorityQueue).(PQStruct)
		fmt.Printf("(%d, %d) ", item.value, item.priority)
	}
	fmt.Println()
}
