package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack struct {
	myStack list.List
	min     []int
}

func Constructor155() MinStack {
	return MinStack{list.List{}, []int{math.MaxInt64}}
}

func (ms *MinStack) Push(x int) {
	ms.myStack.PushBack(x)
	if x <= ms.min[len(ms.min)-1] {
		ms.min = append(ms.min, x)
	}
}

func (ms *MinStack) Pop() {
	tmp := ms.myStack.Back().Value.(int)
	ms.myStack.Remove(ms.myStack.Back())
	if tmp == ms.min[len(ms.min)-1] {
		ms.min = ms.min[:len(ms.min)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.myStack.Back().Value.(int)
}

func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}

func main() {
	minStack := Constructor155()
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
