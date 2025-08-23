package main

import "fmt"

type ListNode234 struct {
	Val  int
	Next *ListNode234
}

func isPalindrome234(head *ListNode234) bool {
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
	a := ListNode234{1, nil}
	b := ListNode234{2, &a}
	c := ListNode234{2, &b}
	d := ListNode234{1, &c}
	fmt.Println(isPalindrome234(&d))
}
