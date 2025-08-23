package main

import "fmt"

type ListNode61 struct {
	Val  int
	Next *ListNode61
}

func rotateRight(head *ListNode61, k int) *ListNode61 {
	l, prev, ans := 0, head, head
	for prev != nil {
		l++
		prev = prev.Next
	}
	if l < 2 {
		return head
	}
	for i := 0; i < k%l; i++ {
		ans = rotate61(ans)
	}
	return ans
}

func rotate61(head *ListNode61) *ListNode61 {
	l, p1, p2 := 0, head, head
	for p1 != nil {
		l++
		p1 = p1.Next
	}
	for i := 0; i < l-2; i++ {
		p2 = p2.Next
	}
	nxt := p2.Next
	p2.Next = nil
	nxt.Next = head
	return nxt
}

func main() {
	e := &ListNode61{5, nil}
	d := &ListNode61{4, e}
	c := &ListNode61{3, d}
	b := &ListNode61{2, c}
	a := &ListNode61{1, b}
	fmt.Println(rotateRight(a, 3))
}
