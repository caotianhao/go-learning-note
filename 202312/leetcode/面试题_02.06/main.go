package main

import "fmt"

type ListNode0206 struct {
	Val  int
	Next *ListNode0206
}

func isPalindrome0206(head *ListNode0206) bool {
	val := make([]int, 0)
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}
	l := len(val)
	for i := 0; i < l; i++ {
		if val[i] != val[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome0206(nil))
}
