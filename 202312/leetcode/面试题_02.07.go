package main

import "fmt"

type ListNode0207 struct {
	Val  int
	Next *ListNode0207
}

func getIntersectionNode0207(headA, headB *ListNode0207) *ListNode0207 {
	hashMap := map[*ListNode0207]bool{}
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
	fmt.Println(getIntersectionNode0207(nil, nil))
}
