package main

import (
	"fmt"
	"sort"
)

type TreeNode2583 struct {
	Val   int
	Left  *TreeNode2583
	Right *TreeNode2583
}

func kthLargestLevelSum(root *TreeNode2583, k int) int64 {
	// 新模板
	// 不用 list 了，太慢
	// 用正经的切片模拟队列
	// 直接入队，不用判断是否为空
	queue, levelSum := []*TreeNode2583{root}, make([]int, 0)
	for len(queue) > 0 {
		tmp, partSum := queue, 0
		queue = nil
		for _, node := range tmp {
			partSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelSum = append(levelSum, partSum)
	}
	sort.Ints(levelSum)
	if k > len(levelSum) {
		return -1
	}
	return int64(levelSum[len(levelSum)-k])
}

func main() {
	four := &TreeNode2583{4, nil, nil}
	six := &TreeNode2583{6, nil, nil}
	two := &TreeNode2583{2, four, six}
	one := &TreeNode2583{1, nil, nil}
	three := &TreeNode2583{3, nil, nil}
	seven := &TreeNode2583{7, nil, nil}
	eight := &TreeNode2583{8, two, one}
	nine := &TreeNode2583{9, three, seven}
	five := &TreeNode2583{5, eight, nine}
	fmt.Println(kthLargestLevelSum(five, 2))
}
