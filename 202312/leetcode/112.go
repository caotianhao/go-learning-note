package main

import "fmt"

type TreeNode112 struct {
	Val   int
	Left  *TreeNode112
	Right *TreeNode112
}

//func hasPathSum(root *TreeNode112, targetSum int) bool {
//	var dfs func(*TreeNode112, int)
//	flag := false
//	dfs = func(node112 *TreeNode112, sum int) {
//		if node112 != nil {
//			sum += node112.Val
//			if node112.Left == nil && node112.Right == nil {
//				if sum == targetSum {
//					flag = true
//					return
//				}
//			} else {
//				dfs(node112.Left, sum)
//				dfs(node112.Right, sum)
//			}
//		}
//	}
//	dfs(root, 0)
//	return flag
//}

func hasPathSum(root *TreeNode112, targetSum int) (flag bool) {
	var dfs func(*TreeNode112, int)
	dfs = func(node *TreeNode112, s int) {
		if node != nil {
			s -= node.Val
			if node.Left == nil && node.Right == nil {
				if s == 0 {
					flag = true
					return
				}
			} else {
				dfs(node.Left, s)
				dfs(node.Right, s)
			}
		}
	}
	dfs(root, targetSum)
	return
}

func main() {
	fmt.Println(hasPathSum(nil, 112))
}
