package main

import (
	"container/list"
	"fmt"
)

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	var queue list.List
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		for i := 0; i < queue.Len(); i++ {
			t := queue.Front().Value.(*Node116)
			queue.Remove(queue.Front())
			if t.Left != nil {
				queue.PushBack(t.Left)
				t.Left.Next = t.Right
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
				if t.Next == nil {
					t.Right.Next = nil
				} else {
					t.Right.Next = t.Next.Left
				}
			}
		}
	}
	return root
}

func main() {
	seven := &Node116{7, nil, nil, nil}
	six := &Node116{6, nil, nil, nil}
	five := &Node116{5, nil, nil, nil}
	four := &Node116{4, nil, nil, nil}
	three := &Node116{3, six, seven, nil}
	two := &Node116{2, four, five, nil}
	one := &Node116{1, two, three, nil}
	fmt.Println(connect(one))
}
