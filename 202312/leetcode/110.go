package main

import "fmt"

type TreeNode110 struct {
	Val   int
	Left  *TreeNode110
	Right *TreeNode110
}

func isBalanced(root *TreeNode110) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) &&
		isBalanced(root.Right) &&
		abs110(treeHeight(root.Left)-treeHeight(root.Right)) <= 1
}

func treeHeight(root *TreeNode110) int {
	if root == nil {
		return 0
	}
	return max(treeHeight(root.Left), treeHeight(root.Right)) + 1
}

func abs110(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	three := TreeNode110{3, nil, nil}
	two := TreeNode110{2, &three, nil}
	one := TreeNode110{1, nil, &two}
	fmt.Println(isBalanced(&one))
}
