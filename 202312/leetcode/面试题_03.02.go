package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack0302 struct {
	myStack list.List
	min     []int
}

func Constructor0302() MinStack0302 {
	return MinStack0302{list.List{}, []int{math.MaxInt64}}
}

func (ms *MinStack0302) Push(x int) {
	ms.myStack.PushBack(x)
	if x <= ms.min[len(ms.min)-1] {
		ms.min = append(ms.min, x)
	}
}

func (ms *MinStack0302) Pop() {
	tmp := ms.myStack.Back().Value.(int)
	ms.myStack.Remove(ms.myStack.Back())
	if tmp == ms.min[len(ms.min)-1] {
		ms.min = ms.min[:len(ms.min)-1]
	}
}

func (ms *MinStack0302) Top() int {
	return ms.myStack.Back().Value.(int)
}

func (ms *MinStack0302) GetMin() int {
	return ms.min[len(ms.min)-1]
}

func main() {
	minStack := Constructor0302()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Push(-3)
	minStack.Push(-4)
	fmt.Println(minStack.GetMin()) //-4

	minStack.Pop()
	fmt.Println(minStack.Top())    //-3
	fmt.Println(minStack.GetMin()) //-3

	minStack.Pop()
	fmt.Println(minStack.GetMin()) //-3
	minStack.Pop()
	fmt.Println(minStack.GetMin()) //-2
}
