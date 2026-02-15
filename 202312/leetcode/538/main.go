package main

import "fmt"

type TreeNode538 struct {
	Val   int
	Left  *TreeNode538
	Right *TreeNode538
}

func convertBST(root *TreeNode538) *TreeNode538 {
	source := make([]int, 0)
	var dfs func(node *TreeNode538)
	dfs = func(node *TreeNode538) {
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
	var dfs2 func(node *TreeNode538)
	dfs2 = func(node *TreeNode538) {
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
	fmt.Println(convertBST(nil))
}
