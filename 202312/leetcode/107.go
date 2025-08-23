package main

import (
	"container/list"
	"fmt"
)

type TreeNode107 struct {
	Val   int
	Left  *TreeNode107
	Right *TreeNode107
}

func levelOrderBottom(root *TreeNode107) [][]int {
	var queue list.List
	res := make([][]int, 0)
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		a, size := make([]int, 0), queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*TreeNode107)
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
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return res
}

func main() {
	seven := &TreeNode107{7, nil, nil}
	fifteen := &TreeNode107{15, nil, nil}
	twenty := &TreeNode107{20, fifteen, seven}
	nine := &TreeNode107{9, nil, nil}
	three := &TreeNode107{3, nine, twenty}
	fmt.Println(levelOrderBottom(three))
}
