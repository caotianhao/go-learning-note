package main

import "fmt"

type TreeNode965 struct {
	Val   int
	Left  *TreeNode965
	Right *TreeNode965
}

func isUnivalTree(root *TreeNode965) bool {
	tmp := root.Val
	flag := true
	var dfs func(node *TreeNode965)
	dfs = func(node *TreeNode965) {
		if node != nil {
			dfs(node.Left)
			dfs(node.Right)
			if node.Val != tmp {
				flag = false
			}
		}
	}
	dfs(root)
	return flag
}

func main() {
	p := &TreeNode965{}
	fmt.Println(isUnivalTree(p))
}
