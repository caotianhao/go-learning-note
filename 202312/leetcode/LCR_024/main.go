package main

import "fmt"

type ListNode2024 struct {
	Val  int
	Next *ListNode2024
}

func reverseList2024(head *ListNode2024) *ListNode2024 {
	var pre *ListNode2024
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	fmt.Println(reverseList2024(nil))
}
