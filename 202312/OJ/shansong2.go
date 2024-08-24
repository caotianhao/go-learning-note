package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getHeight(node.Left)
		right := getHeight(node.Right)
		if left == -1 || right == -1 || absSS(left-right) > 1 {
			return -1
		}
		return maxSS(left, right) + 1
	}
	return getHeight(root) >= 0
}

func maxSS(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absSS(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isBalanced(nil))
}
