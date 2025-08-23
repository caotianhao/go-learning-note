package main

import (
	"container/list"
	"fmt"
)

type TreeNode1038 struct {
	Val   int
	Left  *TreeNode1038
	Right *TreeNode1038
}

func bstToGst(root *TreeNode1038) *TreeNode1038 {
	source := make([]int, 0)
	var dfs func(node *TreeNode1038)
	dfs = func(node *TreeNode1038) {
		if node != nil {
			dfs(node.Left)
			source = append(source, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	l := len(source)
	need := make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		need[i] = need[i+1] + source[i]
	}
	i := 0
	var dfs2 func(node *TreeNode1038)
	dfs2 = func(node *TreeNode1038) {
		if node != nil {
			dfs2(node.Left)
			node.Val = need[i]
			i++
			dfs2(node.Right)
		}
	}
	dfs2(root)
	return root
}

func main() {
	eight := &TreeNode1038{8, nil, nil}
	three := &TreeNode1038{3, nil, nil}
	zero := &TreeNode1038{0, nil, nil}
	five := &TreeNode1038{5, nil, nil}
	two := &TreeNode1038{2, nil, three}
	seven := &TreeNode1038{7, nil, eight}
	one := &TreeNode1038{1, zero, two}
	six := &TreeNode1038{6, five, seven}
	four := &TreeNode1038{4, one, six}
	t := bstToGst(four)
	arr := make([]int, 0)
	var queue list.List
	queue.PushBack(t)
	for queue.Len() != 0 {
		t := queue.Front().Value.(*TreeNode1038)
		queue.Remove(queue.Front())
		arr = append(arr, t.Val)
		if t.Left != nil {
			queue.PushBack(t.Left)
		}
		if t.Right != nil {
			queue.PushBack(t.Right)
		}
	}
	fmt.Println(arr)
}
