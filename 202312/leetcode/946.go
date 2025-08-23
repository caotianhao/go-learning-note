package main

import (
	"container/list"
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	i, j, l, pushStack := 0, 0, len(pushed), list.List{}
	for ; i < l; i++ {
		pushStack.PushBack(pushed[i])
		for pushStack.Back().Value.(int) == popped[j] {
			pushStack.Remove(pushStack.Back())
			j++
			if j >= l || pushStack.Len() == 0 {
				break
			}
		}
	}
	return pushStack.Len() == 0
}

func main() {
	//fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})) // t
	//fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})) // f
	//fmt.Println(validateStackSequences([]int{1, 2}, []int{2, 1}))                   // t
	//fmt.Println(validateStackSequences([]int{1, 2}, []int{1, 2}))                   // t
	fmt.Println(validateStackSequences([]int{1}, []int{1}))                         // t
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4})) // t
	fmt.Println(validateStackSequences([]int{}, []int{}))                           // t
}
