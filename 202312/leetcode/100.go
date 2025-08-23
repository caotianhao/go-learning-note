package main

import "fmt"

type TreeNode100 struct {
	Val   int
	Left  *TreeNode100
	Right *TreeNode100
}

func isSameTree(p *TreeNode100, q *TreeNode100) bool {
	var dfs func(*TreeNode100, *TreeNode100) bool
	dfs = func(r1, r2 *TreeNode100) bool {
		if r1 == nil && r2 != nil {
			return false
		} else if r1 != nil && r2 == nil {
			return false
		} else if r1 == nil && r2 == nil {
			return true
		} else {
			return r1.Val == r2.Val && dfs(r1.Left, r2.Left) && dfs(r1.Right, r2.Right)
		}
	}
	return dfs(p, q)
}

func main() {
	a3 := &TreeNode100{2, nil, nil}
	a2 := &TreeNode100{3, nil, nil}
	a1 := &TreeNode100{1, a2, a3}
	b3 := &TreeNode100{3, nil, nil}
	b2 := &TreeNode100{2, nil, nil}
	b1 := &TreeNode100{1, b2, b3}
	fmt.Println(isSameTree(a1, b1))
}
