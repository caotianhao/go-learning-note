package main

import "fmt"

type TreeNode236 struct {
	Val   int
	Left  *TreeNode236
	Right *TreeNode236
}

func lowestCommonAncestor(root, p, q *TreeNode236) *TreeNode236 {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func main() {
	p, q := &TreeNode236{}, &TreeNode236{}
	fmt.Println(lowestCommonAncestor(nil, p, q))
}
