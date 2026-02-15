package main

import "fmt"

type TreeNode101 struct {
	Val   int
	Left  *TreeNode101
	Right *TreeNode101
}

func isSymmetric101(root *TreeNode101) bool {
	if root == nil {
		return true
	}
	var check func(node1, node2 *TreeNode101) bool
	check = func(node1, node2 *TreeNode101) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil {
			return false
		}
		if node2 == nil {
			return false
		}
		return node1.Val == node2.Val && check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
	}
	return check(root.Left, root.Right)
}

func main() {
	fmt.Println(isSymmetric101(nil))
}
