package main

import "fmt"

type TreeNodeLCR124 struct {
	Val   int
	Left  *TreeNodeLCR124
	Right *TreeNodeLCR124
}

func deduceTree(preorder []int, inorder []int) *TreeNodeLCR124 {
	if len(preorder) == 0 {
		return nil
	}
	r := -1
	for i, v := range inorder {
		if v == preorder[0] {
			r = i
			break
		}
	}
	root := &TreeNodeLCR124{Val: preorder[0]}
	root.Left = deduceTree(preorder[1:r+1], inorder[:r])
	root.Right = deduceTree(preorder[r+1:], inorder[r+1:])
	return root
}

func main() {
	fmt.Println(deduceTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
