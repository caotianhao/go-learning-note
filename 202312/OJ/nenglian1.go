package main

import "fmt"

type ListNodeNL struct {
	Val  int
	Next *ListNodeNL
}

func partition(pHead *ListNodeNL, x int) *ListNodeNL {
	dummyHead := &ListNodeNL{Next: pHead}
	var click bool
	prev, ll, nxt := dummyHead, dummyHead, dummyHead
	cur := pHead
	for cur != nil {
		var tmp *ListNodeNL
		if cur.Val >= x && !click {
			ll = prev
			nxt = cur
			click = true
		}
		if cur.Val < x && click {
			tmp = cur.Next
			prev.Next = cur.Next
			cur.Next = nxt
			ll.Next = cur
			ll = cur
		}
		if tmp != nil {
			cur = tmp
			continue
		}
		prev = cur
		cur = cur.Next
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(partition(nil, 1))
}
