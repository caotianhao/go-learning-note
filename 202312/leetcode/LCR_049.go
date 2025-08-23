package main

import (
	"fmt"
)

type TreeNode2049 struct {
	Val   int
	Left  *TreeNode2049
	Right *TreeNode2049
}

func sumNumbers2049(root *TreeNode2049) (s int) {
	var dfs func(*TreeNode2049, int)
	dfs = func(node *TreeNode2049, sum int) {
		if node != nil {
			sum = sum*10 + node.Val
			if node.Left == nil && node.Right == nil {
				s += sum
				return
			} else {
				dfs(node.Left, sum)
				dfs(node.Right, sum)
			}
		}
	}
	dfs(root, 0)
	return
}

func main() {
	three := &TreeNode2049{3, nil, nil}
	two := &TreeNode2049{2, nil, nil}
	one := &TreeNode2049{1, two, three}
	fmt.Println(sumNumbers2049(one)) //25
}
