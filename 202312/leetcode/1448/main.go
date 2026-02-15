package main

import (
	"fmt"
	"math"
)

type TreeNode1448 struct {
	Val   int
	Left  *TreeNode1448
	Right *TreeNode1448
}

func goodNodes(root *TreeNode1448) int {
	return dfs1448(root, math.MinInt64)
}

func dfs1448(node *TreeNode1448, cur int) (cnt int) {
	if node == nil {
		return 0
	}
	if node.Val >= cur {
		cnt++
		cur = node.Val
	}
	cnt += dfs1448(node.Left, cur) + dfs1448(node.Right, cur)
	return
}

func main() {
	a := &TreeNode1448{3, nil, nil}
	b := &TreeNode1448{1, nil, nil}
	c := &TreeNode1448{5, nil, nil}
	d := &TreeNode1448{1, a, nil}
	e := &TreeNode1448{4, b, c}
	f := &TreeNode1448{3, d, e}
	fmt.Println(goodNodes(f))
}
