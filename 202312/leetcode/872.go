package main

import "fmt"

type TreeNode872 struct {
	Val   int
	Left  *TreeNode872
	Right *TreeNode872
}

func leafSimilar(root1, root2 *TreeNode872) bool {
	var dfs func(node *TreeNode872)
	var arr []int
	dfs = func(node *TreeNode872) {
		if node != nil {
			if node.Left == nil && node.Right == nil {
				arr = append(arr, node.Val)
			}
			dfs(node.Left)
			dfs(node.Right)
		}
		return
	}
	dfs(root1)
	arr1 := arr
	arr = []int{}
	dfs(root2)
	arr2 := arr
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	four := &TreeNode872{4, nil, nil}
	seven := &TreeNode872{7, nil, nil}
	six := &TreeNode872{6, nil, nil}
	nine := &TreeNode872{9, nil, nil}
	eight := &TreeNode872{8, nil, nil}
	two := &TreeNode872{2, seven, four}
	five := &TreeNode872{5, six, two}
	one := &TreeNode872{1, nine, eight}
	three := &TreeNode872{3, five, one}
	three1 := &TreeNode872{3, one, five}
	fmt.Println(leafSimilar(three, three1))
}
