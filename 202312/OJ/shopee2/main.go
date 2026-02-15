package main

import (
	"fmt"
	"sort"
)

type TreeNodeShopee struct {
	Val   int
	Left  *TreeNodeShopee
	Right *TreeNodeShopee
}

func getMostGold(node *TreeNodeShopee) []int {
	maxGold := maxPath(getAllRoute(node))
	q := pathMax(node, maxGold)
	return q[0]
}

func pathMax(root *TreeNodeShopee, maxGold int) [][]int {
	path := make([]int, 0)
	var dfs func(*TreeNodeShopee, int)
	res := make([][]int, 0)
	dfs = func(node *TreeNodeShopee, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && left == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, maxGold)
	return res
}

func getAllRoute(root *TreeNodeShopee) (arr []int) {
	var dfs func(*TreeNodeShopee, int)
	dfs = func(node *TreeNodeShopee, s int) {
		if node.Right == nil && node.Left == nil {
			arr = append(arr, s)
		}
		if node.Left != nil {
			dfs(node.Left, s+node.Left.Val)
		}
		if node.Right != nil {
			dfs(node.Right, s+node.Right.Val)
		}
	}
	dfs(root, root.Val)
	return
}

func maxPath(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func main() {
	a := &TreeNodeShopee{2, nil, nil}
	b := &TreeNodeShopee{4, a, nil}
	c := &TreeNodeShopee{6, nil, nil}
	d := &TreeNodeShopee{2, b, nil}
	e := &TreeNodeShopee{3, nil, c}
	r := &TreeNodeShopee{1, d, e}
	fmt.Println(getMostGold(r))
}
