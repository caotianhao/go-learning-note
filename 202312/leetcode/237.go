package main

import "fmt"

type ListNode237 struct {
	Val  int
	Next *ListNode237
}

func deleteNode237(node *ListNode237) {
	node.Val = node.Next.Val
	// 之前的困惑：题目只说了node不是倒数第一个节点
	// 但是假设是倒数第二个节点呢，那么它还有.Next.Next吗
	// 答案是有的，因为nil也可以赋值给其他节点
	// 这样就好理解了
	// 就是先把node.Next的值给node
	// 然后不管node.Next.Next是节点还是nil都直接赋值给node.Next就可以了
	node.Next = node.Next.Next
}

func main() {
	p := &ListNode237{6, nil}
	q := &ListNode237{5, p}
	a := &ListNode237{4, q}
	b := &ListNode237{3, a}
	c := &ListNode237{2, b}
	d := &ListNode237{1, c}
	e, f := d, d

	for e != nil {
		fmt.Print(e.Val, " ")
		e = e.Next
	}

	fmt.Println()
	deleteNode237(b)

	for f != nil {
		fmt.Print(f.Val, " ")
		f = f.Next
	}
}
