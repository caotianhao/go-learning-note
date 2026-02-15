package main

import (
	"fmt"

	"github.com/Arafatk/DataViz/trees/redblacktree"
)

type TreeNode230 struct {
	Val   int
	Left  *TreeNode230
	Right *TreeNode230
}

func kthSmallest(root *TreeNode230, k int) int {
	rbt := redblacktree.NewWithIntComparator()
	var dfs func(node *TreeNode230)
	dfs = func(node *TreeNode230) {
		if node != nil {
			dfs(node.Left)
			rbt.Put(node.Val, nil)
			if rbt.Size() > k {
				rbt.Remove(rbt.Right().Key)
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return rbt.Right().Key.(int)
}

func main() {
	two := &TreeNode230{2, nil, nil}
	one := &TreeNode230{1, nil, two}
	four := &TreeNode230{4, nil, nil}
	three := &TreeNode230{3, one, four}
	fmt.Println(kthSmallest(three, 1))
}
