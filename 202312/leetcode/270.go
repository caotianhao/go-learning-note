package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode270 struct {
	Val   int
	Left  *TreeNode270
	Right *TreeNode270
}

func closestValue(root *TreeNode270, target float64) int {
	minN := math.MaxFloat64
	res := make([]int, 0)
	var dfs func(node *TreeNode270)
	var dfs1 func(node *TreeNode270)
	dfs = func(node *TreeNode270) {
		if node != nil {
			dfs(node.Left)
			if math.Abs(float64(node.Val)-target) < minN {
				minN = math.Abs(float64(node.Val) - target)
			}
			dfs(node.Right)
		}
	}
	dfs1 = func(node *TreeNode270) {
		if node != nil {
			dfs1(node.Left)
			if math.Abs(float64(node.Val)-target) == minN {
				res = append(res, node.Val)
			}
			dfs1(node.Right)
		}
	}
	dfs(root)
	dfs1(root)
	sort.Ints(res)
	return res[0]
}

func main() {
	one := &TreeNode270{1, nil, nil}
	three := &TreeNode270{3, nil, nil}
	two := &TreeNode270{2, one, three}
	five := &TreeNode270{5, nil, nil}
	four := &TreeNode270{4, two, five}
	fmt.Println(closestValue(four, 4.5))
}
