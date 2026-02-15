package main

import "fmt"

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	dummy := &ListNode19{0, head}
	d, h := dummy, head
	for i := 0; i < n; i++ {
		h = h.Next
	}
	for ; h != nil; h = h.Next {
		d = d.Next
	}
	d.Next = d.Next.Next
	return dummy.Next
}

func main() {
	a := ListNode19{5, nil}
	b := ListNode19{4, &a}
	c := ListNode19{3, &b}
	d := ListNode19{2, &c}
	e := ListNode19{1, &d}
	fmt.Println(removeNthFromEnd(&e, 2))
}
