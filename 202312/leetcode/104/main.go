package main

import "fmt"

type TreeNode104 struct {
	Val   int
	Left  *TreeNode104
	Right *TreeNode104
}

func maxDepth104(root *TreeNode104) int {
	if root == nil {
		return 0
	}
	return max(maxDepth104(root.Left), maxDepth104(root.Right)) + 1
}

func main() {
	three := TreeNode104{3, nil, nil}
	two := TreeNode104{2, &three, nil}
	one := TreeNode104{1, nil, &two}
	fmt.Println(maxDepth104(&one))
}
