package main

import "fmt"

type TreeNode2045 struct {
	Val   int
	Left  *TreeNode2045
	Right *TreeNode2045
}

func findBottomLeftValue2045(root *TreeNode2045) int {
	queue := []*TreeNode2045{root}
	res := -1
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		res = node.Val
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return res
}

func main() {
	c := &TreeNode2045{3, nil, nil}
	b := &TreeNode2045{2, nil, nil}
	a := &TreeNode2045{1, c, b}
	fmt.Println(findBottomLeftValue2045(a))
}
