package main

import "fmt"

type ListNodeLCR142 struct {
	Val  int
	Next *ListNodeLCR142
}

func trainningPlan(l1, l2 *ListNodeLCR142) *ListNodeLCR142 {
	dummy := &ListNodeLCR142{}
	var cur = dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return dummy.Next
}

func main() {
	l11 := &ListNodeLCR142{4, nil}
	l12 := &ListNodeLCR142{2, l11}
	l1 := &ListNodeLCR142{1, l12}
	l21 := &ListNodeLCR142{4, nil}
	l22 := &ListNodeLCR142{3, l21}
	l2 := &ListNodeLCR142{9, l22}
	p := trainningPlan(l1, l2)
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println()
	fmt.Println(trainningPlan(nil, nil))
	fmt.Println(trainningPlan(l1, nil))
	fmt.Println(trainningPlan(nil, l2))
}
