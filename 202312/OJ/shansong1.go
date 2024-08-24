package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(nil))
}
