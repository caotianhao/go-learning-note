package main

import "fmt"

type ListNode92 struct {
	Val  int
	Next *ListNode92
}

func reverseBetween(head *ListNode92, left, right int) *ListNode92 {
	// dummy
	dummy := &ListNode92{Next: head}

	// beforeLeft: left前面的节点
	// rightNode: right位置的节点
	// leftNode: left位置的节点
	// afterRight: right后面的节点
	var beforeLeft, rightNode, leftNode, afterRight *ListNode92

	// 找到 left 前面的节点
	beforeLeft = dummy
	for i := 0; i < left-1; i++ {
		beforeLeft = beforeLeft.Next
	}

	// 找到对应 right 的节点
	rightNode = beforeLeft
	for i := 0; i < right+1-left; i++ {
		rightNode = rightNode.Next
	}

	// 找到剩余两个节点
	leftNode, afterRight = beforeLeft.Next, rightNode.Next

	// 断开连接
	beforeLeft.Next, rightNode.Next = nil, nil

	// 反转
	reverse92(leftNode)

	// 接回
	beforeLeft.Next, leftNode.Next = rightNode, afterRight

	return dummy.Next
}

func reverse92(head *ListNode92) {
	var pre *ListNode92
	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = pre
		pre = cur
		cur = n
	}
}

func main() {
	e := &ListNode92{55, nil}
	d := &ListNode92{44, e}
	c := &ListNode92{33, d}
	b := &ListNode92{22, c}
	a := &ListNode92{11, b}
	fmt.Println(reverseBetween(a, 2, 4))
}
