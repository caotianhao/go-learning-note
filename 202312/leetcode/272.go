package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode272 struct {
	Val   int
	Left  *TreeNode272
	Right *TreeNode272
}

func closestKValues(root *TreeNode272, target float64, k int) (res []int) {
	var dfs func(node *TreeNode272)
	m := make([][]float64, 0)
	dfs = func(node *TreeNode272) {
		if node != nil {
			dfs(node.Left)
			m = append(m, []float64{float64(node.Val), math.Abs(float64(node.Val) - target)})
			dfs(node.Right)
		}
	}
	dfs(root)
	sort.Slice(m, func(i, j int) bool {
		return m[i][1] < m[j][1]
	})
	for i := 0; i < k; i++ {
		res = append(res, int(m[i][0]))
	}
	return
}

func main() {
	three := &TreeNode272{3, nil, nil}
	one := &TreeNode272{1, nil, nil}
	two := &TreeNode272{2, one, three}
	five := &TreeNode272{5, nil, nil}
	four := &TreeNode272{4, two, five}
	fmt.Println(closestKValues(four, 3.7, 2))
}
