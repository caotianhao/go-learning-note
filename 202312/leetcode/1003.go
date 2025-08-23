package main

import (
	"container/list"
	"fmt"
)

func isValid1003(s string) bool {
	var stack list.List
	for _, v := range s {
		stack.PushBack(v)
		if stack.Len() >= 3 && stack.Back().Value.(int32) == 'c' &&
			stack.Back().Prev().Value.(int32) == 'b' &&
			stack.Back().Prev().Prev().Value.(int32) == 'a' {
			stack.Remove(stack.Back())
			stack.Remove(stack.Back())
			stack.Remove(stack.Back())
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid1003("abcbcbcbabcbacbac"))
}
