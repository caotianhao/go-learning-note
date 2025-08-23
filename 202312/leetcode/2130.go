package main

import "fmt"

type ListNode2130 struct {
	Val  int
	Next *ListNode2130
}

func pairSum(head *ListNode2130) (maxN int) {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		maxN = max(maxN, arr[left]+arr[right])
		left++
		right--
	}
	return
}

func main() {
	fmt.Println(pairSum(nil))
}
