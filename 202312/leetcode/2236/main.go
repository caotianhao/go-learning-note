package main

import "fmt"

type TreeNode2236 struct {
	Val   int
	Left  *TreeNode2236
	Right *TreeNode2236
}

func checkTree(root *TreeNode2236) bool {
	return root.Val == root.Right.Val+root.Left.Val
}

func main() {
	ll := &TreeNode2236{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	rr := &TreeNode2236{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode2236{
		Val:   12,
		Left:  ll,
		Right: rr,
	}
	fmt.Println(checkTree(&root))
}
