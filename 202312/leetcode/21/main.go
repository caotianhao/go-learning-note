package main

import (
	"fmt"
)

type ListNode21 struct {
	Val  int
	Next *ListNode21
}

//func mergeTwoLists(list1 *ListNode21, list2 *ListNode21) *ListNode21 {
//	arr := make([]int, 0)
//	for list1 != nil {
//		arr = append(arr, list1.Val)
//		list1 = list1.Next
//	}
//	for list2 != nil {
//		arr = append(arr, list2.Val)
//		list2 = list2.Next
//	}
//	pre := &ListNode21{}
//	head := &ListNode21{}
//	pre.Next = head
//	sort.Ints(arr)
//	for i := 0; i < len(arr); i++ {
//		head.Next = &ListNode21{arr[i], nil}
//		head = head.Next
//	}
//	return pre.Next.Next
//}

func mergeTwoLists(list1 *ListNode21, list2 *ListNode21) *ListNode21 {
	head := &ListNode21{-1, nil}
	pre := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	} else {
		pre.Next = list1
	}
	return head.Next
}

func main() {
	a4 := ListNode21{4, nil}
	a2 := ListNode21{2, &a4}
	list1 := ListNode21{1, &a2}

	b4 := ListNode21{4, nil}
	b3 := ListNode21{3, &b4}
	list2 := ListNode21{1, &b3}

	newList := mergeTwoLists(&list1, &list2)
	fmt.Println("nl", newList)
	for newList != nil {
		fmt.Print(newList.Val)
		fmt.Print(" ")
		newList = newList.Next
	}
}
