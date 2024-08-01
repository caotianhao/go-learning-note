package main

import "fmt"

const capacity = 10

type MyStack struct {
	Stack [capacity]int
	Cnt   int
}

func (s *MyStack) Push(v int) {
	if s.Cnt < capacity {
		s.Stack[s.Cnt] = v
		s.Cnt++
		return
	} else {
		fmt.Println("the Stack is full")
	}
}

func (s *MyStack) Pop() (v int) {
	if s.Cnt != 0 {
		v = s.Stack[s.Cnt-1]
		s.Cnt--
		return
	} else {
		fmt.Println("the Stack is empty")
		return 0
	}
}

func (s *MyStack) String() (r string) {
	for i := 0; i < s.Cnt; i++ {
		r += fmt.Sprintf("%c", s.Stack[i])
	}
	return
}

func main() {
	st := &MyStack{}
	st.Push(65)
	st.Push(66)
	st.Push(67)
	fmt.Println(st.Pop())
	fmt.Println(st.String())
}
