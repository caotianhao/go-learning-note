package main

import "fmt"

type TreeNode226 struct {
	Val   int
	Left  *TreeNode226
	Right *TreeNode226
}

func invertTree226(root *TreeNode226) *TreeNode226 {
	if root != nil {
		l, r := root.Left, root.Right
		root.Left, root.Right = r, l
		invertTree226(root.Left)
		invertTree226(root.Right)
	}
	return root
}

func main() {
	three := TreeNode226{3, nil, nil}
	two := TreeNode226{2, &three, nil}
	one := TreeNode226{1, nil, &two}
	fmt.Println(invertTree226(&one))
}
