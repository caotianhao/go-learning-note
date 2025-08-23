package main

import "fmt"

type TreeNode2265 struct {
	Val   int
	Left  *TreeNode2265
	Right *TreeNode2265
}

func averageOfSubtree(root *TreeNode2265) (cnt int) {
	var dfs func(node *TreeNode2265)
	var judge func(node *TreeNode2265) bool
	judge = func(node *TreeNode2265) bool {
		sum, cnt := 0, 0
		var dfs func(n *TreeNode2265)
		dfs = func(n *TreeNode2265) {
			if n != nil {
				sum += n.Val
				cnt++
				dfs(n.Left)
				dfs(n.Right)
			}
		}
		dfs(node)
		return node.Val == sum/cnt
	}
	dfs = func(node *TreeNode2265) {
		if node != nil {
			if judge(node) {
				cnt++
			}
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	fmt.Println(averageOfSubtree(nil))
}
