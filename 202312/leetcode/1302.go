package main

import (
	"container/list"
	"fmt"
)

type TreeNode1302 struct {
	Val   int
	Left  *TreeNode1302
	Right *TreeNode1302
}

func deepestLeavesSum(root *TreeNode1302) int {
	var queue list.List
	if root != nil {
		queue.PushBack(root)
	}
	arr := make([][]int, 0)
	for queue.Len() != 0 {
		tmp, size := make([]int, 0), queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*TreeNode1302)
			queue.Remove(queue.Front())
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				queue.PushBack(t.Left)
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
			}
		}
		arr = append(arr, tmp)
	}
	sum := 0
	for _, v := range arr[len(arr)-1] {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(deepestLeavesSum(nil))
}
