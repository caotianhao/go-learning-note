package main

import "fmt"

type TreeNode897 struct {
	Val   int
	Left  *TreeNode897
	Right *TreeNode897
}

func increasingBST(root *TreeNode897) *TreeNode897 {
	var dfs func(node *TreeNode897)
	a := &TreeNode897{}
	res := a
	dfs = func(node *TreeNode897) {
		if node != nil {
			dfs(node.Left)
			tmp := &TreeNode897{node.Val, nil, nil}
			res.Right = tmp
			res = tmp
			dfs(node.Right)
		}
	}
	dfs(root)
	return a.Right
}

func main() {
	fmt.Println(increasingBST(nil))
}
