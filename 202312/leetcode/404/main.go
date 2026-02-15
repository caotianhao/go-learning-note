package main

import "fmt"

type TreeNode404 struct {
	Val   int
	Left  *TreeNode404
	Right *TreeNode404
}

func sumOfLeftLeaves(root *TreeNode404) int {
	sum := 0
	var dfs func(node *TreeNode404)
	dfs = func(node *TreeNode404) {
		if node != nil {
			dfs(node.Left)
			dfs(node.Right)
			if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			}
		}
	}
	dfs(root)
	return sum
}

func main() {
	seven := &TreeNode404{7, nil, nil}
	fifteen := &TreeNode404{15, nil, nil}
	twenty := &TreeNode404{20, fifteen, seven}
	nine := &TreeNode404{9, nil, nil}
	three := &TreeNode404{3, nine, twenty}
	fmt.Println(sumOfLeftLeaves(three))
}
