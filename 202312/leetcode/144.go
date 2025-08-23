package main

import "fmt"

type TreeNode144 struct {
	Val   int
	Left  *TreeNode144
	Right *TreeNode144
}

//func preorderTraversal(root *TreeNode144) (res []int) {
//	var truePreorderTraversal func(node *TreeNode144)
//	truePreorderTraversal = func(node *TreeNode144) {
//		if node != nil {
//			res = append(res, node.Val)
//			truePreorderTraversal(node.Left)
//			truePreorderTraversal(node.Right)
//		}
//	}
//	truePreorderTraversal(root)
//	return
//}

func preorderTraversal(root *TreeNode144) (res []int) {
	var dfs func(node *TreeNode144)
	dfs = func(node *TreeNode144) {
		if node != nil {
			res = append(res, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	three := TreeNode144{3, nil, nil}
	two := TreeNode144{2, &three, nil}
	one := TreeNode144{1, nil, &two}
	fmt.Println(preorderTraversal(&one))
}
