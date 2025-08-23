package main

import (
	"container/list"
	"fmt"
)

type Node429 struct {
	Val      int
	Children []*Node429
}

func levelOrder429(root *Node429) [][]int {
	var queue list.List
	res := make([][]int, 0)
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		tmp, size := make([]int, 0), queue.Len()
		for i := 0; i < size; i++ {
			t := queue.Front().Value.(*Node429)
			queue.Remove(queue.Front())
			tmp = append(tmp, t.Val)
			for j := 0; j < len(t.Children); j++ {
				queue.PushBack(t.Children[j])
			}
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
	fmt.Println(levelOrder429(nil))
}
