package main

import (
	"fmt"

	"github.com/Arafatk/DataViz/trees/redblacktree"
)

type TreeNode671 struct {
	Val   int
	Left  *TreeNode671
	Right *TreeNode671
}

func findSecondMinimumValue(root *TreeNode671) int {
	rbt := redblacktree.NewWithIntComparator()
	var dfs func(node *TreeNode671)
	dfs = func(node *TreeNode671) {
		if node != nil {
			rbt.Put(node.Val, nil)
			if rbt.Size() >= 2 {
				rbt.Remove(rbt.Right().Key)
			}
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	if rbt.Size() == 2 {
		return rbt.Right().Key.(int)
	}
	return -1
}

func main() {
	fmt.Println(findSecondMinimumValue(nil))
}
