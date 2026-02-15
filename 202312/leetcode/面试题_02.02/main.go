package main

import "fmt"

type ListNode0202 struct {
	Val  int
	Next *ListNode0202
}

func kthToLast(head *ListNode0202, k int) int {
	slow, fast := head, head
	for cnt := 0; cnt < k; cnt++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

func main() {
	five := &ListNode0202{5, nil}
	four := &ListNode0202{4, five}
	three := &ListNode0202{3, four}
	two := &ListNode0202{2, three}
	one := &ListNode0202{1, two}
	fmt.Println(kthToLast(one, 2))
}
