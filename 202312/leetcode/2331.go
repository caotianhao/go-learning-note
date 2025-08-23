package main

import "fmt"

type TreeNode2331 struct {
	Val   int
	Left  *TreeNode2331
	Right *TreeNode2331
}

//func evaluateTree(root *TreeNode2331) bool {
//	for !isLeaf(root) {
//		if isLeaf(root.Left) && isLeaf(root.Right) {
//			if root.Val == 2 {
//				if root.Left.Val == 0 && root.Right.Val == 0 {
//					root.Val = 0
//				} else {
//					root.Val = 1
//				}
//			} else {
//				if root.Left.Val == 1 && root.Right.Val == 1 {
//					root.Val = 1
//				} else {
//					root.Val = 0
//				}
//			}
//			root.Left, root.Right = nil, nil
//		} else if isLeaf(root.Left) && !isLeaf(root.Right) {
//			evaluateTree(root.Right)
//		} else if !isLeaf(root.Left) && isLeaf(root.Right) {
//			evaluateTree(root.Left)
//		} else if !isLeaf(root.Left) && !isLeaf(root.Right) {
//			evaluateTree(root.Left)
//			evaluateTree(root.Right)
//		}
//	}
//	if root.Val == 0 {
//		return false
//	}
//	return true
//}
//
//func isLeaf(root *TreeNode2331) bool {
//	return root.Left == nil && root.Right == nil
//}

func evaluateTree(root *TreeNode2331) bool {
	if root.Left == nil {
		return root.Val == 1
	} else if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}

func main() {
	// 0 false  1 true  2 or  3 and
	t1 := TreeNode2331{0, nil, nil}
	t2 := TreeNode2331{1, nil, nil}
	t3 := TreeNode2331{3, &t1, &t2}
	t4 := TreeNode2331{1, nil, nil}
	t5 := TreeNode2331{2, &t4, &t3}
	fmt.Println(evaluateTree(&t5)) //true
}
