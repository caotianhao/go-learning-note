package main

import (
	"sync/atomic"
)

type LockFreeQueue struct {
	queue []interface{}
	len   int32
	head  int32
	tail  int32
}

func NewQueue(n int32) *LockFreeQueue {
	q := &LockFreeQueue{
		queue: make([]interface{}, n+1, n+1),
		len:   n + 1,
	}
	return q
}

func (s *LockFreeQueue) EnQueue(v interface{}) {
	for {
		tail := atomic.LoadInt32(&s.tail)
		n := (tail + 1) % s.len
		if atomic.CompareAndSwapInt32(&s.head, n, n) {
			continue // 队列满
		}
		if !atomic.CompareAndSwapInt32(&s.tail, tail, n) {
			continue // 获取失败
		}
		s.queue[tail] = v
		break
	}
}

func (s *LockFreeQueue) DeQueue() interface{} {
	for {
		tail := atomic.LoadInt32(&s.tail)
		head := atomic.LoadInt32(&s.head)
		if tail == head {
			continue
		}
		n := (head + 1) % s.len
		if !atomic.CompareAndSwapInt32(&s.head, head, n) {
			continue
		}
		return s.queue[head]
	}
}
