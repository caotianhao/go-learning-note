package main

import "fmt"

type TreeNode235 struct {
	Val   int
	Left  *TreeNode235
	Right *TreeNode235
}

func lowestCommonAncestor235(root, p, q *TreeNode235) *TreeNode235 {
	res := root
	for {
		if res.Val > p.Val && res.Val > q.Val {
			res = res.Left
		} else if res.Val < p.Val && res.Val < q.Val {
			res = res.Right
		} else {
			return res
		}
	}
}

func main() {
	a := &TreeNode235{15, nil, nil}
	b := &TreeNode235{22, nil, nil}
	c := &TreeNode235{18, a, b}
	fmt.Println(lowestCommonAncestor235(c, a, b))
}
