package main

import "fmt"

type ListNode876 struct {
	Val  int
	Next *ListNode876
}

//func middleNode(head *ListNode876) *ListNode876 {
//	cnt, cnt1 := 0, 1
//	for tmp := head.Next; tmp != nil; tmp = tmp.Next {
//		cnt++
//	}
//	if cnt == 0 {
//		return head
//	}
//	if cnt%2 == 0 {
//		for tmp := head.Next; ; tmp = tmp.Next {
//			cnt1++
//			if cnt1 == cnt/2+1 {
//				return tmp
//			}
//		}
//	} else {
//		for tmp := head.Next; ; tmp = tmp.Next {
//			if cnt1 == cnt/2+1 {
//				return tmp
//			}
//			cnt1++
//		}
//	}
//}

func middleNode(head *ListNode876) *ListNode876 {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	a := ListNode876{6, nil}
	aa := ListNode876{5, &a}
	b := ListNode876{4, &aa}
	c := ListNode876{3, &b}
	d := ListNode876{2, &c}
	e := ListNode876{1, &d}
	fmt.Println(middleNode(&e))
}
