package main

import "fmt"

type TreeNodeTx struct {
	Val   int
	Left  *TreeNodeTx
	Right *TreeNodeTx
}

func pathNumber(root *TreeNodeTx) (cnt int) {
	p := pathNumber2(root)
	for _, v := range p {
		c := 0
		for _, vv := range v {
			if vv == 1 {
				c++
			} else {
				c--
			}
		}
		if c > 0 {
			cnt++
		}
	}
	return
}

func pathNumber2(root *TreeNodeTx) [][]int {
	var dfs func(*TreeNodeTx)
	path := make([][]int, 0)
	currPath := make([]int, 0)
	dfs = func(node *TreeNodeTx) {
		if node == nil {
			return
		}
		currPath = append(currPath, node.Val)
		if node.Left == nil && node.Right == nil {
			path = append(path, append([]int(nil), currPath...))
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		currPath = currPath[:len(currPath)-1]
	}
	dfs(root)
	return path
}

func main() {
	fmt.Println(pathNumber(nil))
}
