package main

import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) (head *ListNode2) {
	//创建尾结点
	var tail *ListNode2
	//进位
	tmp := 0
	//当有链表不空的时候循环
	for l1 != nil || l2 != nil {
		//传值，如果有长度不相等的，那么长度短的自然默认为0
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//计算和
		sum := n1 + n2 + tmp
		//计算进位并处理
		sum, tmp = sum%10, sum/10
		if head == nil {
			//首尾指针的用法
			//如果头结点为空，则在头结点处创建结点，首尾指向同一个结点
			head = &ListNode2{Val: sum}
			tail = head
		} else {
			//如果头结点不为空，在尾结点之后创建新结点并移动尾指针
			if tail == nil {
				fmt.Println("tail is nil")
				return
			}
			tail.Next = &ListNode2{Val: sum}
			tail = tail.Next
		}
	}
	//最后如果有大进位，则需要再创建一个新的结点并进 1
	if tmp == 1 && tail != nil {
		tail.Next = &ListNode2{Val: 1}
	}
	return
}

func main() {
	//243,564
	n12 := ListNode2{3, nil}
	n11 := ListNode2{4, &n12}
	n10 := ListNode2{2, &n11}

	n22 := ListNode2{4, nil}
	n21 := ListNode2{6, &n22}
	n20 := ListNode2{5, &n21}

	arr := make([]int, 0)
	i := addTwoNumbers(&n10, &n20)
	for i != nil {
		arr = append(arr, i.Val)
		i = i.Next
	}
	fmt.Println(arr)
}
