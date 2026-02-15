package main

import "fmt"

type TreeNode145 struct {
	Val   int
	Left  *TreeNode145
	Right *TreeNode145
}

//func postorderTraversal(root *TreeNode145) (res []int) {
//	var truePostorderTraversal func(node *TreeNode145)
//	truePostorderTraversal = func(node *TreeNode145) {
//		if node != nil {
//			truePostorderTraversal(node.Left)
//			truePostorderTraversal(node.Right)
//			res = append(res, node.Val)
//		}
//	}
//	truePostorderTraversal(root)
//	return
//}

func postorderTraversal(root *TreeNode145) (res []int) {
	var dfs func(node *TreeNode145)
	dfs = func(node *TreeNode145) {
		if node != nil {
			dfs(node.Left)
			dfs(node.Right)
			res = append(res, node.Val)
		}
	}
	dfs(root)
	return
}

func main() {
	three := TreeNode145{3, nil, nil}
	two := TreeNode145{2, &three, nil}
	one := TreeNode145{1, nil, &two}
	fmt.Println(postorderTraversal(&one))
}
