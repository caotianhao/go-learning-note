package main

import (
	"container/list"
	"fmt"
)

type TreeNode637 struct {
	Val   int
	Left  *TreeNode637
	Right *TreeNode637
}

func averageOfLevels(root *TreeNode637) (res []float64) {
	var queue list.List
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		tmpArr, sum, size := make([]int, 0), 0, queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*TreeNode637)
			queue.Remove(queue.Front())
			tmpArr = append(tmpArr, t.Val)
			if t.Left != nil {
				queue.PushBack(t.Left)
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
			}
		}
		for i := 0; i < len(tmpArr); i++ {
			sum += tmpArr[i]
		}
		res = append(res, float64(sum)/float64(len(tmpArr)))
	}
	return
}

func main() {
	seven := &TreeNode637{7, nil, nil}
	fifteen := &TreeNode637{15, nil, nil}
	twenty := &TreeNode637{20, fifteen, seven}
	nine := &TreeNode637{9, nil, nil}
	three := &TreeNode637{3, nine, twenty}
	fmt.Println(averageOfLevels(three))
}
