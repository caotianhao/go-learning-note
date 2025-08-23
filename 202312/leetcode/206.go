package main

import "fmt"

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

//func reverseList(head *ListNode206) *ListNode206 {
//	arr := make([]int, 0)
//	for head != nil {
//		arr = append(arr, head.Val)
//		head = head.Next
//	}
//	l := len(arr)
//	for i := 0; i < l/2; i++ {
//		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
//	}
//	tr := &ListNode206{}
//	c := tr
//	for i := 0; i < l; i++ {
//		tmp := &ListNode206{arr[i], nil}
//		c.Next = tmp
//		c = c.Next
//	}
//	return tr.Next
//}

//func reverseList(head *ListNode206) *ListNode206 {
//	// 创建一个空节点
//	// 不能用 pre := &ListNode206{}
//	var pre *ListNode206
//	// 常规操作
//	cur := head
//	for cur != nil {
//		// 先记录当前节点的下一个
//		next := cur.Next
//		// 翻转当前节点
//		cur.Next = pre
//		// 继续直到 cur == nil
//		pre = cur
//		cur = next
//	}
//	// 注意返回的是pre而不是cur，因为pre=cur之后cur=next了
//	return pre
//}

func reverseList(head *ListNode206) *ListNode206 {
	var pre *ListNode206
	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = pre
		pre = cur
		cur = n
	}
	return pre
}

func main() {
	a := &ListNode206{5, nil}
	b := &ListNode206{4, a}
	c := &ListNode206{3, b}
	d := &ListNode206{2, c}
	e := &ListNode206{1, d}
	fmt.Println(reverseList(e))
}
