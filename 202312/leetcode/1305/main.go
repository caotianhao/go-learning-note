package main

import (
	"fmt"
	"sort"
)

type TreeNode1305 struct {
	Val   int
	Left  *TreeNode1305
	Right *TreeNode1305
}

func getAllElements(root1, root2 *TreeNode1305) (arr []int) {
	var dfs func(*TreeNode1305)
	dfs = func(node *TreeNode1305) {
		if node != nil {
			dfs(node.Left)
			arr = append(arr, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root1)
	dfs(root2)
	sort.Ints(arr)
	return
}

func main() {
	fmt.Println(getAllElements(&TreeNode1305{22, nil, nil}, &TreeNode1305{7, nil, nil}))
}
