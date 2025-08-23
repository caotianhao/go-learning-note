package main

import "fmt"

type TreeNode513 struct {
	Val   int
	Left  *TreeNode513
	Right *TreeNode513
}

// 请找出该二叉树的 最底层 最左边 节点的值
// 思路就是从右往左层序遍历
// 最后一个节点即为所求
func findBottomLeftValue(root *TreeNode513) int {
	queue := []*TreeNode513{root}
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
	c := &TreeNode513{3, nil, nil}
	b := &TreeNode513{2, nil, nil}
	a := &TreeNode513{1, b, c}
	fmt.Println(findBottomLeftValue(a))
}
