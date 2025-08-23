package main

import "fmt"

type ListNode24 struct {
	Val  int
	Next *ListNode24
}

func swapPairs(head *ListNode24) *ListNode24 {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i := 0; i < len(arr)-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	fmt.Println(arr)
	newHead := &ListNode24{}
	c := newHead
	for i := 0; i < len(arr); i++ {
		tmp := &ListNode24{arr[i], nil}
		c.Next = tmp
		c = c.Next
	}
	return newHead.Next
}

func main() {
	//aa := &ListNode24{5, nil}
	a := &ListNode24{4, nil}
	b := &ListNode24{3, a}
	c := &ListNode24{2, b}
	d := &ListNode24{1, c}
	fmt.Println(swapPairs(d))
}
