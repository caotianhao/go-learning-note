package main

import (
	"fmt"
)

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

func inorderTraversal(root *TreeNode94) (res []int) {

	var dfs func(node *TreeNode94)
	dfs = func(node *TreeNode94) {
		if node != nil {
			dfs(node.Left)
			res = append(res, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	three := TreeNode94{3, nil, nil}
	two := TreeNode94{2, &three, nil}
	one := TreeNode94{1, nil, &two}
	fmt.Println(inorderTraversal(&one))

	//nn := TreeNode94{}
	//fmt.Println(inorderTraversal(&nn))
	//
	//one1 := TreeNode94{1, nil, nil}
	//fmt.Println(inorderTraversal(&one1))
}
