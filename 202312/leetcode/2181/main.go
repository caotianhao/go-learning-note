package main

import "fmt"

type ListNode2181 struct {
	Val  int
	Next *ListNode2181
}

func mergeNodes(head *ListNode2181) *ListNode2181 {
	dummy := &ListNode2181{}
	newHead := dummy
	sum := 0
	for head != nil {
		if head.Val == 0 {
			tmp := &ListNode2181{sum, nil}
			newHead.Next = tmp
			newHead = newHead.Next
			sum = 0
		} else {
			sum += head.Val
		}
		head = head.Next
	}
	return dummy.Next.Next
}

func main() {
	a := &ListNode2181{0, nil}
	b := &ListNode2181{67, a}
	c := &ListNode2181{40, b}
	d := &ListNode2181{0, c}
	e := &ListNode2181{4, d}
	f := &ListNode2181{0, e}
	t := mergeNodes(f)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}
