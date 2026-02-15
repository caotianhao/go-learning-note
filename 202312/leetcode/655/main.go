package main

import (
	"fmt"
	"strconv"
)

type TreeNode655 struct {
	Val   int
	Left  *TreeNode655
	Right *TreeNode655
}

func getHeight655(root *TreeNode655) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight655(root.Left), getHeight655(root.Right))
}

func printTree(root *TreeNode655) [][]string {
	height := getHeight655(root)
	m := height
	n := (1 << m) - 1
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = ""
		}
	}
	var dfs func(*TreeNode655, int, int)
	dfs = func(node *TreeNode655, r, c int) {
		if node != nil {
			res[r][c] = strconv.Itoa(node.Val)
			if node.Left != nil {
				dfs(node.Left, r+1, c-(1<<(height-r-2)))
			}
			if node.Right != nil {
				dfs(node.Right, r+1, c+(1<<(height-r-2)))
			}
		}
	}
	dfs(root, 0, (n-1)/2)
	return res
}

func main() {
	a := &TreeNode655{2, nil, nil}
	b := &TreeNode655{1, a, nil}
	fmt.Println(printTree(b))

	four := &TreeNode655{4, nil, nil}
	two := &TreeNode655{2, nil, four}
	three := &TreeNode655{3, nil, nil}
	one := &TreeNode655{1, two, three}
	fmt.Println(printTree(one))

	fmt.Println(printTree(nil))
}
