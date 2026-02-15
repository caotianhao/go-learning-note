package main

import (
	"fmt"
	"math"
)

type TreeNode1161 struct {
	Val   int
	Left  *TreeNode1161
	Right *TreeNode1161
}

func maxLevelSum(root *TreeNode1161) (r int) {
	queue := []*TreeNode1161{root}
	level, maxN := 1, math.MinInt64
	for len(queue) != 0 {
		t, s := queue, 0
		queue = nil
		for _, node := range t {
			s += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if s > maxN {
			maxN = s
			r = level
		}
		level++
	}
	return
}

func main() {
	fmt.Println(maxLevelSum(&TreeNode1161{18, nil, nil}))
}
