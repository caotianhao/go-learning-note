package main

import (
	"container/list"
	"fmt"
)

type SortedStack struct {
	mainStack list.List
}

var tempStack list.List

func Constructor0305() SortedStack {
	return SortedStack{list.List{}}
}

func (ss *SortedStack) Push(val int) {
	if ss.IsEmpty() || val <= ss.mainStack.Back().Value.(int) {
		ss.mainStack.PushBack(val)
		return
	}
	if !ss.IsEmpty() && val > ss.mainStack.Back().Value.(int) {
		for val > ss.mainStack.Back().Value.(int) {
			tempStack.PushBack(ss.mainStack.Back().Value.(int))
			ss.mainStack.Remove(ss.mainStack.Back())
			if ss.IsEmpty() {
				break
			}
		}
		ss.mainStack.PushBack(val)
		for tempStack.Len() != 0 {
			ss.mainStack.PushBack(tempStack.Back().Value.(int))
			tempStack.Remove(tempStack.Back())
		}
	}
}

func (ss *SortedStack) Pop() {
	if ss.IsEmpty() {
		return
	}
	ss.mainStack.Remove(ss.mainStack.Back())
}

func (ss *SortedStack) Peek() int {
	if ss.IsEmpty() {
		return -1
	}
	return ss.mainStack.Back().Value.(int)
}

func (ss *SortedStack) IsEmpty() bool {
	return ss.mainStack.Len() == 0
}

func main() {
	s := Constructor0305()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
}
