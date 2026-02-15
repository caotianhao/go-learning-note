package main

import "fmt"

type ListNode023 struct {
	Val  int
	Next *ListNode023
}

func getIntersectionNode023(headA, headB *ListNode023) *ListNode023 {
	hm := make(map[*ListNode023]struct{})
	for headA != nil {
		hm[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := hm[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	fmt.Println(getIntersectionNode023(nil, nil))
}
