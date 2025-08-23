package main

import "fmt"

type TreeNode501 struct {
	Val   int
	Left  *TreeNode501
	Right *TreeNode501
}

func findMode(root *TreeNode501) (res []int) {
	// base当前数 count当前数的数量 max最大count也就是众数
	base, count, maxN := 0, 0, 0
	// 统计众数的函数
	var findArr func(x int)
	// 中序遍历
	var dfs func(node *TreeNode501)
	// 中序遍历的二叉搜索树一定是非递减排序的，这一点很重要
	findArr = func(x int) {
		// 从头遍历统计众数，思想很巧妙
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxN {
			res = append(res, x)
		} else if count > maxN {
			maxN = count
			res = []int{base}
		}
	}
	dfs = func(node *TreeNode501) {
		if node != nil {
			dfs(node.Left)
			findArr(node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	two1 := TreeNode501{2, nil, nil}
	two2 := TreeNode501{2, &two1, nil}
	one := TreeNode501{1, nil, &two2}
	fmt.Println(findMode(&one))
}
