package main

import (
	"fmt"
	"strconv"
)

type TreeNode606 struct {
	Val   int
	Left  *TreeNode606
	Right *TreeNode606
}

func tree2str(root *TreeNode606) (str string) {
	if root == nil {
		return
	}
	str += strconv.FormatInt(int64(root.Val), 10)
	if root.Left != nil && root.Right == nil {
		str += "(" + tree2str(root.Left) + ")"
	} else if root.Left == nil && root.Right != nil {
		str += "()(" + tree2str(root.Right) + ")"
	} else if root.Left != nil && root.Right != nil {
		str += "(" + tree2str(root.Left) + ")(" + tree2str(root.Right) + ")"
	}
	return
}

func main() {
	four := TreeNode606{4, nil, nil}
	three := TreeNode606{3, nil, nil}
	two := TreeNode606{2, nil, &four}
	//two := TreeNode606{2, &four, nil}
	one := TreeNode606{1, &two, &three}
	fmt.Println(tree2str(&one))
	fmt.Println(tree2str(nil))
	fmt.Println(tree2str(&four))
}
