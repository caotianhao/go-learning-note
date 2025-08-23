package main

import "fmt"

type ListNode143 struct {
	Val  int
	Next *ListNode143
}

func reorderList(head *ListNode143) {
	mid := getMid143(head)
	h1 := head
	h2 := mid.Next
	mid.Next = nil
	linkList143(h1, reverse143(h2))
}

func getMid143(head *ListNode143) *ListNode143 {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse143(head *ListNode143) *ListNode143 {
	var p *ListNode143
	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = p
		p = cur
		cur = n
	}
	return p
}

func linkList143(h1, h2 *ListNode143) {
	h := &ListNode143{}
	pre := h
	flag := 1
	for h1 != nil && h2 != nil {
		if flag == 1 {
			pre.Next = h1
			h1 = h1.Next
			pre = pre.Next
			flag = 1 - flag
		} else {
			pre.Next = h2
			h2 = h2.Next
			pre = pre.Next
			flag = 1 - flag
		}
	}
	if h1 != nil {
		pre.Next = h1
	}
	if h2 != nil {
		pre.Next = h2
	}
	//h1 = h.Next
	//var t1, t2 *ListNode143
	//for h1 != nil && h2 != nil {
	//	t1 = h1.Next
	//	t2 = h2.Next
	//
	//	h1.Next = h2
	//	h1 = t1
	//
	//	h2.Next = h1
	//	h2 = t2
	//}
}

func main() {
	a := &ListNode143{5, nil}
	b := &ListNode143{4, a}
	c := &ListNode143{3, b}
	d := &ListNode143{2, c}
	e := &ListNode143{1, d}
	fmt.Println(e.Next.Val)
	reorderList(e)
	fmt.Println(e.Next.Val)
}
