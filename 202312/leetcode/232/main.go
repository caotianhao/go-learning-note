package main

import (
	"container/list"
	"fmt"
)

//type MyQueue struct {
//	arr []int
//}
//
//func Constructor232() MyQueue {
//	return MyQueue{[]int{}}
//}
//
//func (q *MyQueue) Push(x int) {
//	q.arr = append(q.arr, x)
//}
//
//func (q *MyQueue) Pop() int {
//	tmp := q.arr[0]
//	q.arr = q.arr[1:]
//	return tmp
//}
//
//func (q *MyQueue) Peek() int {
//	return q.arr[0]
//}
//
//func (q *MyQueue) Empty() bool {
//	return len(q.arr) == 0
//}

type MyQueue struct {
	q list.List
}

func Constructor232() MyQueue {
	queue := list.List{}
	return MyQueue{queue}
}

func (mq *MyQueue) Push(x int) {
	mq.q.PushBack(x)
}

func (mq *MyQueue) Pop() int {
	tmp := mq.q.Front().Value.(int)
	mq.q.Remove(mq.q.Front())
	return tmp
}

func (mq *MyQueue) Peek() int {
	tmp := mq.q.Front().Value.(int)
	return tmp
}

func (mq *MyQueue) Empty() bool {
	return mq.q.Len() == 0
}

func main() {
	queue := Constructor232()
	queue.Push(1)
	//queue.Push(2)
	//fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
