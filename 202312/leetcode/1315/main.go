package main

import "fmt"

type TreeNode1315 struct {
	Val   int
	Left  *TreeNode1315
	Right *TreeNode1315
}

func sumEvenGrandparent(root *TreeNode1315) (sum int) {
	var dfs = func(node *TreeNode1315) {}
	dfs = func(node *TreeNode1315) {
		if node != nil {
			if node.Val&1 == 0 {
				if node.Left != nil {
					if node.Left.Left != nil {
						sum += node.Left.Left.Val
					}
					if node.Left.Right != nil {
						sum += node.Left.Right.Val
					}
				}
				if node.Right != nil {
					if node.Right.Left != nil {
						sum += node.Right.Left.Val
					}
					if node.Right.Right != nil {
						sum += node.Right.Right.Val
					}
				}
			}
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	fmt.Println(sumEvenGrandparent(nil))
}
