package main

import "fmt"

type ListNode1019 struct {
	Val  int
	Next *ListNode1019
}

func nextLargerNodes(head *ListNode1019) (res []int) {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l := len(arr)
	for i := 0; i < l; i++ {
		flag := true
		for j := 0; i+j < l; j++ {
			if arr[i+j] > arr[i] {
				res = append(res, arr[i+j])
				flag = false
				break
			}
		}
		if flag {
			res = append(res, 0)
		}
	}
	return
}

func main() {
	a := ListNode1019{5, nil}
	b := ListNode1019{1, &a}
	c := ListNode1019{2, &b}
	fmt.Println(nextLargerNodes(&c)) // 5,5,0
	d := ListNode1019{5, nil}
	e := ListNode1019{3, &d}
	f := ListNode1019{4, &e}
	g := ListNode1019{7, &f}
	h := ListNode1019{2, &g}
	fmt.Println(nextLargerNodes(&h)) // 7,0,5,5,0
	fmt.Println(nextLargerNodes(&d)) // 7,0,5,5,0
}
