package main

import "fmt"

type TreeNode2053 struct {
	Val   int
	Left  *TreeNode2053
	Right *TreeNode2053
}

func inorderSuccessor2053(root, p *TreeNode2053) (r *TreeNode2053) {
	var dfs func(node *TreeNode2053)
	arr, index, target := make([]int, 0), -1, -1
	dfs = func(node *TreeNode2053) {
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
	var dfs1 func(node *TreeNode2053)
	dfs1 = func(node *TreeNode2053) {
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
	p := &TreeNode2053{}
	fmt.Println(inorderSuccessor2053(nil, p))
	//fmt.Println(isucc)
}
