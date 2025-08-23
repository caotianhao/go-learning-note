package main

import (
	"container/list"
	"fmt"
)

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

func levelOrder(root *TreeNode102) (res [][]int) {
	var queue list.List
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		a, size := make([]int, 0), queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*TreeNode102)
			queue.Remove(queue.Front())
			a = append(a, t.Val)
			if t.Left != nil {
				queue.PushBack(t.Left)
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
			}
		}
		res = append(res, a)
	}
	return
}

func levelOrderTest(root *TreeNode102) (res []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode102{root}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		res = append(res, t.Val)
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
	return
}

func main() {
	seven := &TreeNode102{7, nil, nil}
	fifteen := &TreeNode102{15, nil, nil}
	twenty := &TreeNode102{20, fifteen, seven}
	nine := &TreeNode102{9, nil, nil}
	three := &TreeNode102{3, nine, twenty}
	fmt.Println(levelOrder(three))
	fmt.Println(levelOrderTest(three))
}
