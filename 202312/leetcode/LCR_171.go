package main

import "fmt"

type ListNodeLCR171 struct {
	Val  int
	Next *ListNodeLCR171
}

func getIntersectionNode171(headA, headB *ListNodeLCR171) *ListNodeLCR171 {
	hm := make(map[*ListNodeLCR171]struct{})
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
	fmt.Println(getIntersectionNode171(nil, nil))
}
