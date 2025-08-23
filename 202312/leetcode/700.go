package main

import "fmt"

type TreeNode700 struct {
	Val   int
	Left  *TreeNode700
	Right *TreeNode700
}

func searchBST(root *TreeNode700, val int) *TreeNode700 {
	//要先判空，否则直接取 Val 要是它空就报错了
	//if root.Val == val || root == nil {
	//if root == nil || root.Val == val {
	//	return root
	//} else if root.Val > val {
	//	return searchBST(root.Left, val)
	//}
	//return searchBST(root.Right, val)

	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func main() {
	three := TreeNode700{3, nil, nil}
	one := TreeNode700{1, nil, nil}
	two := TreeNode700{2, &one, &three}
	seven := TreeNode700{7, nil, nil}
	four := TreeNode700{4, &two, &seven}
	fmt.Println(three, one, two, seven, four)
	fmt.Println(searchBST(&four, 5))
}
