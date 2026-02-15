package main

import "fmt"

type ListNode2807 struct {
	Val  int
	Next *ListNode2807
}

func insertGreatestCommonDivisors(head *ListNode2807) *ListNode2807 {
	if head == nil || head.Next == nil {
		return head
	}
	c := head
	for c.Next != nil {
		nxt := c.Next
		v := getGCD2807(c.Val, c.Next.Val)
		t := &ListNode2807{Val: v}
		c.Next = t
		t.Next = nxt
		c = nxt
	}
	return head
}

func getGCD2807(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD2807(b, a%b)
}

func main() {
	//fmt.Println(getGCD2807(18, 6))
	fmt.Println(insertGreatestCommonDivisors(nil))
}
