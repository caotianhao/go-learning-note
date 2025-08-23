package main

import "fmt"

type TreeNode1026 struct {
	Val   int
	Left  *TreeNode1026
	Right *TreeNode1026
}

func maxAncestorDiff(root *TreeNode1026) int {
	var dfs func(node *TreeNode1026, min, max int) int
	dfs = func(node *TreeNode1026, minN, maxN int) int {
		if node == nil {
			return 0
		}
		diff := max(abs1026(node.Val-minN), abs1026(node.Val-maxN))
		minN, maxN = min(minN, node.Val), max(maxN, node.Val)
		diff = max(diff, dfs(node.Left, minN, maxN))
		diff = max(diff, dfs(node.Right, minN, maxN))
		return diff
	}
	return dfs(root, root.Val, root.Val)
}

func abs1026(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	a4 := &TreeNode1026{4, nil, nil}
	a7 := &TreeNode1026{7, nil, nil}
	a13 := &TreeNode1026{13, nil, nil}
	a1 := &TreeNode1026{1, nil, nil}
	a6 := &TreeNode1026{6, a4, a7}
	a14 := &TreeNode1026{14, a13, nil}
	a3 := &TreeNode1026{3, a1, a6}
	a10 := &TreeNode1026{10, nil, a14}
	a8 := &TreeNode1026{8, a3, a10}
	fmt.Println(maxAncestorDiff(a8))
}
