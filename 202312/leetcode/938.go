package main

import "fmt"

type TreeNode938 struct {
	Val   int
	Left  *TreeNode938
	Right *TreeNode938
}

func rangeSumBST(root *TreeNode938, low, high int) (sum int) {
	var dfs func(node *TreeNode938)
	dfs = func(node *TreeNode938) {
		if node != nil {
			dfs(node.Left)
			if node.Val >= low && node.Val <= high {
				sum += node.Val
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	fmt.Println(rangeSumBST(nil, 0, 0))
}
