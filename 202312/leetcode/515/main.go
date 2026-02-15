package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode515 struct {
	Val   int
	Left  *TreeNode515
	Right *TreeNode515
}

func largestValues(root *TreeNode515) (res []int) {
	var queue list.List
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		tmp, size := math.MinInt64, queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*TreeNode515)
			queue.Remove(queue.Front())
			if t.Val > tmp {
				tmp = t.Val
			}
			if t.Left != nil {
				queue.PushBack(t.Left)
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
			}
		}
		res = append(res, tmp)
	}
	return
}

func main() {
	fmt.Println(largestValues(nil))
}
