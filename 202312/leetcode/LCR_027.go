package main

import "fmt"

type ListNode2027 struct {
	Val  int
	Next *ListNode2027
}

func isPalindrome2027(head *ListNode2027) bool {
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
	fmt.Println(isPalindrome2027(nil))
}
