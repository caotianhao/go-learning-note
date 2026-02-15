package main

import (
	"fmt"
	"strconv"
)

type TreeNode1022 struct {
	Val   int
	Left  *TreeNode1022
	Right *TreeNode1022
}

func sumRootToLeaf(root *TreeNode1022) int {
	sum := 0
	var dfs func(node *TreeNode1022, path string)
	dfs = func(node *TreeNode1022, path string) {
		if node != nil {
			path += strconv.Itoa(node.Val)
			if node.Left == nil && node.Right == nil {
				s, _ := strconv.ParseInt(path, 2, 64)
				sum += int(s)
				return
			} else {
				dfs(node.Left, path+"")
				dfs(node.Right, path+"")
			}
		}
	}
	dfs(root, "")
	return sum
}

func main() {
	fmt.Println(sumRootToLeaf(nil))
}
