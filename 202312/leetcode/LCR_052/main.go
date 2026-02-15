package main

import "fmt"

type TreeNode2052 struct {
	Val   int
	Left  *TreeNode2052
	Right *TreeNode2052
}

func increasingBST2052(root *TreeNode2052) *TreeNode2052 {
	var dfs func(node *TreeNode2052)
	a := &TreeNode2052{}
	res := a
	dfs = func(node *TreeNode2052) {
		if node != nil {
			dfs(node.Left)
			tmp := &TreeNode2052{node.Val, nil, nil}
			res.Right = tmp
			res = tmp
			dfs(node.Right)
		}
	}
	dfs(root)
	return a.Right
}

func main() {
	fmt.Println(increasingBST2052(nil))
}
