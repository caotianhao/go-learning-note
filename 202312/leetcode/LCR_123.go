package main

import "fmt"

type ListNodeLCR123 struct {
	Val  int
	Next *ListNodeLCR123
}

func reverseBookList(head *ListNodeLCR123) (arr []int) {
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return
}

func main() {
	a := &ListNodeLCR123{1, nil}
	b := &ListNodeLCR123{221, a}
	c := &ListNodeLCR123{321, b}
	d := &ListNodeLCR123{11243, c}
	fmt.Println(reverseBookList(d))
	fmt.Println(reverseBookList(nil))
	fmt.Println(reverseBookList(a))
}
