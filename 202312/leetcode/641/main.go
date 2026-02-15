package main

import (
	"container/list"
	"fmt"
)

type MyCircularDeque struct {
	queue list.List
	size  int
}

func Constructor641(k int) MyCircularDeque {
	myList := list.List{}
	return MyCircularDeque{myList, k}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue.PushFront(value)
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue.PushBack(value)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue.Remove(q.queue.Front())
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue.Remove(q.queue.Back())
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue.Front().Value.(int)
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue.Back().Value.(int)
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.queue.Len() == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.queue.Len() == q.size
}

func main() {
	c := Constructor641(12)
	fmt.Println(c.IsFull())
}
