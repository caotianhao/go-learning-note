package main

import (
	"fmt"
)

type CQueue struct {
	queue []int
}

func ConstructorLCR125() CQueue {
	return CQueue{[]int{}}
}

func (cq *CQueue) AppendTail(value int) {
	cq.queue = append(cq.queue, value)
}

func (cq *CQueue) DeleteHead() int {
	if len(cq.queue) != 0 {
		o := cq.queue[0]
		cq.queue = cq.queue[1:]
		return o
	}
	return -1
}

func main() {
	q := ConstructorLCR125()
	q.AppendTail(1)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
}
