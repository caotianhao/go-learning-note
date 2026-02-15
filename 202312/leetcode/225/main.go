package main

import (
	"container/list"
	"fmt"
)

//type MyStack struct {
//	arr []int
//}
//
//func Constructor225() MyStack {
//	return MyStack{[]int{}}
//}
//
//func (s *MyStack) Push(x int) {
//	s.arr = append(s.arr, x)
//}
//
//func (s *MyStack) Pop() int {
//	if !s.Empty() {
//		tmp := s.arr[len(s.arr)-1]
//		s.arr = s.arr[:len(s.arr)-1]
//		return tmp
//	} else {
//		return 0
//	}
//}
//
//func (s *MyStack) Top() int {
//	if !s.Empty() {
//		return s.arr[len(s.arr)-1]
//	} else {
//		return 0
//	}
//}
//
//func (s *MyStack) Empty() bool {
//	return len(s.arr) == 0
//}

type MyStack struct {
	st list.List
}

func Constructor225() MyStack {
	stack := list.List{}
	return MyStack{stack}
}

func (s *MyStack) Push(x int) {
	s.st.PushBack(x)
}

func (s *MyStack) Pop() int {
	a := s.st.Back().Value.(int)
	s.st.Remove(s.st.Back())
	return a
}

func (s *MyStack) Top() int {
	a := s.st.Back().Value.(int)
	return a
}

func (s *MyStack) Empty() bool {
	return s.st.Len() == 0
}

func main() {
	stack := Constructor225()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
