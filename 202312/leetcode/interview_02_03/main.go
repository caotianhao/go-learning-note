package main

import "fmt"

type ListNode0203 struct {
	Val  int
	Next *ListNode0203
}

func deleteNode(node *ListNode0203) {
	temp := node.Next
	node.Next = temp.Next
	node.Val = temp.Val
	temp = nil
}

func main() {
	d := ListNode0203{4, nil}
	c := ListNode0203{3, &d}
	b := ListNode0203{2, &c}
	a := ListNode0203{1, &b}
	fmt.Println(a.Next.Val)
	deleteNode(&b)
	fmt.Println(a.Next.Val)
}
