package main

import "fmt"

type TreeNode1712 struct {
	Val   int
	Left  *TreeNode1712
	Right *TreeNode1712
}

func convertBiNode(root *TreeNode1712) *TreeNode1712 {
	var dfs func(node *TreeNode1712)
	a := &TreeNode1712{}
	res := a
	dfs = func(node *TreeNode1712) {
		if node != nil {
			dfs(node.Left)
			tmp := &TreeNode1712{node.Val, nil, nil}
			res.Right = tmp
			res = tmp
			dfs(node.Right)
		}
	}
	dfs(root)
	return a.Right
}

func main() {
	fmt.Println(convertBiNode(nil))
}
