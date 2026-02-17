package main

import "fmt"

type TreeNode0406 struct {
	Val   int
	Left  *TreeNode0406
	Right *TreeNode0406
}

func inorderSuccessor0406(root, p *TreeNode0406) (r *TreeNode0406) {
	var dfs func(node *TreeNode0406)
	arr, index, target := make([]int, 0), -1, -1
	dfs = func(node *TreeNode0406) {
		if node != nil {
			dfs(node.Left)
			arr = append(arr, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	for i := range arr {
		if arr[i] == p.Val {
			index = i
		}
	}
	if index == len(arr)-1 {
		return nil
	} else {
		target = arr[index+1]
	}
	var dfs1 func(node *TreeNode0406)
	dfs1 = func(node *TreeNode0406) {
		if node != nil {
			dfs1(node.Left)
			if node.Val == target {
				r = node
			}
			dfs1(node.Right)
		}
	}
	dfs1(root)
	return r
}

func main() {
	p := &TreeNode0406{}
	fmt.Println(inorderSuccessor0406(nil, p))
}
