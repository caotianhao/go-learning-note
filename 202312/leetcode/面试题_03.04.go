package main

import (
	"container/list"
	"fmt"
)

type MyQueue0304 struct {
	q list.List
}

func Constructor0304() MyQueue0304 {
	queue := list.List{}
	return MyQueue0304{queue}
}

func (mq *MyQueue0304) Push(x int) {
	mq.q.PushBack(x)
}

func (mq *MyQueue0304) Pop() int {
	tmp := mq.q.Front().Value.(int)
	mq.q.Remove(mq.q.Front())
	return tmp
}

func (mq *MyQueue0304) Peek() int {
	return mq.q.Front().Value.(int)
}

func (mq *MyQueue0304) Empty() bool {
	return mq.q.Len() == 0
}

func main() {
	queue := Constructor0304()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
