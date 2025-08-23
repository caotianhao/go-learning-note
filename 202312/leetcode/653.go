package main

import "fmt"

type TreeNode653 struct {
	Val   int
	Left  *TreeNode653
	Right *TreeNode653
}

func findTarget(root *TreeNode653, k int) bool {
	hashMap := map[int]struct{}{}
	ans := false
	var dfs func(node *TreeNode653)
	dfs = func(node *TreeNode653) {
		if node != nil {
			dfs(node.Left)
			if _, ok := hashMap[k-node.Val]; ok {
				ans = true
				return
			}
			hashMap[node.Val] = struct{}{}
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

func main() {
	fmt.Println(findTarget(nil, 11))
}
