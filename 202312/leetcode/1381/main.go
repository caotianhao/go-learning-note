package main

import "fmt"

type CustomStack struct {
	arr     []int
	maxSize int
}

func Constructor1381(maxSize int) CustomStack {
	return CustomStack{[]int{}, maxSize}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.arr) < cs.maxSize {
		cs.arr = append(cs.arr, x)
	}
}

func (cs *CustomStack) Pop() int {
	l := len(cs.arr)
	if l > 0 {
		tmp := cs.arr[l-1]
		cs.arr = cs.arr[:l-1]
		return tmp
	}
	return -1
}

func (cs *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(k, len(cs.arr)); i++ {
		cs.arr[i] += val
	}
}

func main() {
	stk := Constructor1381(3)
	stk.Push(1)            // 栈变为 [1]
	stk.Push(2)            // 栈变为 [1, 2]
	fmt.Println(stk.Pop()) // 返回 2 --> 返回栈顶值 2，栈变为 [1]
	stk.Push(2)            // 栈变为 [1, 2]
	stk.Push(3)            // 栈变为 [1, 2, 3]
	stk.Push(4)            // 栈仍然是 [1, 2, 3]，不能添加其他元素使栈大小变为 4
	stk.Increment(5, 100)  // 栈变为 [101, 102, 103]
	stk.Increment(2, 100)  // 栈变为 [201, 202, 103]
	fmt.Println(stk.Pop()) // 返回 -1 --> 栈为空，返回 -1
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
}
