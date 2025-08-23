package main

import "fmt"

type ListNode203 struct {
	Val  int
	Next *ListNode203
}

func removeElements(head *ListNode203, val int) *ListNode203 {
	dummy := &ListNode203{Next: head}
	for i := dummy; i.Next != nil; {
		if i.Next.Val == val {
			i.Next = i.Next.Next
		} else {
			i = i.Next
		}
	}
	return dummy.Next
}

func main() {
	a := ListNode203{6, nil}
	b := ListNode203{6, &a}
	c := ListNode203{6, &b}
	d := ListNode203{6, &c}
	e := ListNode203{6, &d}
	f := ListNode203{6, &e}
	head := ListNode203{6, &f}
	fmt.Println(removeElements(&head, 6))
}
