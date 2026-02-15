package main

import (
	"fmt"
)

type TreeNode129 struct {
	Val   int
	Left  *TreeNode129
	Right *TreeNode129
}

func sumNumbers(root *TreeNode129) (s int) {
	var dfs func(*TreeNode129, int)
	dfs = func(node *TreeNode129, sum int) {
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
	three := &TreeNode129{3, nil, nil}
	two := &TreeNode129{2, nil, nil}
	one := &TreeNode129{1, two, three}
	fmt.Println(sumNumbers(one)) //25
}
