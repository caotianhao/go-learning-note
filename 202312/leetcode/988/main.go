package main

import (
	"fmt"
)

type TreeNode988 struct {
	Val   int
	Left  *TreeNode988
	Right *TreeNode988
}

func smallestFromLeaf(root *TreeNode988) string {
	var dfs func(*TreeNode988, string)
	minF := "{"
	dfs = func(node *TreeNode988, s string) {
		if node != nil {
			s += string(byte(node.Val + 97))
			if node.Left == nil && node.Right == nil {
				t := reverse988(s)
				if t < minF {
					minF = t
				}
				return
			} else {
				dfs(node.Left, s)
				dfs(node.Right, s)
			}
		}
	}
	dfs(root, "")
	return minF
}

func reverse988(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func main() {
	a := &TreeNode988{3, nil, nil}
	b := &TreeNode988{4, nil, nil}
	c := &TreeNode988{3, nil, nil}
	d := &TreeNode988{4, nil, nil}
	e := &TreeNode988{1, a, b}
	f := &TreeNode988{2, c, d}
	g := &TreeNode988{0, e, f}
	fmt.Println(smallestFromLeaf(g))
}
