package main

import "fmt"

type TreeNode285 struct {
	Val   int
	Left  *TreeNode285
	Right *TreeNode285
}

func inorderSuccessor(root, p *TreeNode285) (r *TreeNode285) {
	var dfs func(node *TreeNode285)
	arr, index, target := make([]int, 0), -1, -1
	dfs = func(node *TreeNode285) {
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
	var dfs1 func(node *TreeNode285)
	dfs1 = func(node *TreeNode285) {
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
	p := &TreeNode285{}
	fmt.Println(inorderSuccessor(nil, p))
}
