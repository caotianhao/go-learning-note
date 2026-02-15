package main

import (
	"fmt"
	"strconv"
)

type TreeNode257 struct {
	Val   int
	Left  *TreeNode257
	Right *TreeNode257
}

func binaryTreePaths(root *TreeNode257) (res []string) {
	// 二叉树路径类入门题
	// 首先 dfs 肯定不能像以前一样只带一个节点参数
	// 需要根据题意带
	// 比如求路径和就带一个之前走过节点的和
	// 本题求路径，就带个分隔符
	var dfs func(*TreeNode257, string)
	dfs = func(node *TreeNode257, de string) {
		// 首先还是先判断根节点是否为空
		if node != nil {
			// 任何路径，皆由根起
			de += strconv.Itoa(node.Val)
			// 若该节点为叶子节点，就把路径加入到路径切片中
			if node.Left == nil && node.Right == nil {
				res = append(res, de)
				// 这个 return 加了效率更高，不加也行
				return
			} else {
				// 若非叶子节点，那就串联之前走过的节点
				// 所以这里也说明了必须有第二个参数
				dfs(node.Left, de+"->")
				dfs(node.Right, de+"->")
			}
		}
	}
	// 调用 dfs
	dfs(root, "")
	// 其实模板熟悉了就很好做路径类的题
	return
}

func main() {
	five := &TreeNode257{5, nil, nil}
	three := &TreeNode257{3, nil, nil}
	two := &TreeNode257{2, nil, five}
	one := &TreeNode257{1, two, three}
	fmt.Println(binaryTreePaths(one))
}
