package main

import "fmt"

type ListNode160 struct {
	Val  int
	Next *ListNode160
}

func getIntersectionNode(headA, headB *ListNode160) *ListNode160 {
	hashMap := map[*ListNode160]bool{}
	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if hashMap[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	a := &ListNode160{5, nil}
	b := &ListNode160{4, a}
	c := &ListNode160{3, b}
	d := &ListNode160{2, c}
	e := &ListNode160{1, d}

	f := &ListNode160{6, b}
	g := &ListNode160{7, f}
	h := &ListNode160{8, g}
	fmt.Println(getIntersectionNode(h, e)) // 4
}
