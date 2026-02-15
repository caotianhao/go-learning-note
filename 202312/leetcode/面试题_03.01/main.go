package main

import (
	"container/list"
	"fmt"
)

type TripleInOne struct {
	stack1, stack2, stack3 list.List
	size                   int
}

func Constructor0301(stackSize int) TripleInOne {
	return TripleInOne{list.List{}, list.List{}, list.List{}, stackSize}
}

func (tio *TripleInOne) Push(stackNum int, value int) {
	if stackNum == 1 {
		if tio.stack1.Len() < tio.size {
			tio.stack1.PushBack(value)
		}
	} else if stackNum == 2 {
		if tio.stack2.Len() < tio.size {
			tio.stack2.PushBack(value)
		}
	} else {
		if tio.stack3.Len() < tio.size {
			tio.stack3.PushBack(value)
		}
	}
}

func (tio *TripleInOne) Pop(stackNum int) int {
	if stackNum == 1 {
		if tio.stack1.Len() != 0 {
			t := tio.stack1.Back().Value.(int)
			tio.stack1.Remove(tio.stack1.Back())
			return t
		}
	} else if stackNum == 2 {
		if tio.stack2.Len() != 0 {
			t := tio.stack2.Back().Value.(int)
			tio.stack2.Remove(tio.stack2.Back())
			return t
		}
	} else {
		if tio.stack3.Len() != 0 {
			t := tio.stack3.Back().Value.(int)
			tio.stack3.Remove(tio.stack3.Back())
			return t
		}
	}
	return -1
}

func (tio *TripleInOne) Peek(stackNum int) int {
	if stackNum == 1 {
		if tio.stack1.Len() != 0 {
			return tio.stack1.Back().Value.(int)
		}
	} else if stackNum == 2 {
		if tio.stack2.Len() != 0 {
			return tio.stack2.Back().Value.(int)
		}
	} else {
		if tio.stack3.Len() != 0 {
			return tio.stack3.Back().Value.(int)
		}
	}
	return -1
}

func (tio *TripleInOne) IsEmpty(stackNum int) bool {
	if stackNum == 1 {
		return tio.stack1.Len() == 0
	} else if stackNum == 2 {
		return tio.stack2.Len() == 0
	} else {
		return tio.stack3.Len() == 0
	}
}

func main() {
	s := Constructor0301(5)
	fmt.Println(s)
}
