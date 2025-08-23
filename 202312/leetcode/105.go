package main

import "fmt"

type TreeNode105 struct {
	Val   int
	Left  *TreeNode105
	Right *TreeNode105
}

func buildTree(preorder, inorder []int) *TreeNode105 {
	if len(preorder) == 0 {
		return nil
	}
	rootIndex := -1
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	root := &TreeNode105{preorder[0], nil, nil}
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])
	return root
}

func main() {
	t := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	// verify
	if t == nil {
		fmt.Println("tree is nil")
		return
	}
	fmt.Println(t.Val, t.Left.Val, t.Right.Val, t.Right.Left.Val, t.Right.Right.Val) //3,9,20,15,7
	//fmt.Println(t.Left.Left)
}
