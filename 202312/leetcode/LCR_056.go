package main

import "fmt"

type TreeNode2056 struct {
	Val   int
	Left  *TreeNode2056
	Right *TreeNode2056
}

func findTarget2056(root *TreeNode2056, k int) bool {
	hashMap := map[int]struct{}{}
	ans := false
	var dfs func(node *TreeNode2056)
	dfs = func(node *TreeNode2056) {
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
	fmt.Println(findTarget2056(nil, 34))
}
