package main

import "fmt"

type TreeNode113 struct {
	Val   int
	Left  *TreeNode113
	Right *TreeNode113
}

//func pathSum(root *TreeNode113, targetSum int) (result [][]int) {
//	var dfs func(*TreeNode113, int, []int)
//	dfs = func(node *TreeNode113, partSum int, path []int) {
//		if node != nil {
//			path = append(path, node.Val)
//			partSum += node.Val
//			if node.Left == nil && node.Right == nil {
//				if partSum == targetSum {
//					result = append(result, append([]int{}, path...))
//				}
//				return
//			} else {
//				dfs(node.Left, partSum, path)
//				dfs(node.Right, partSum, path)
//			}
//		}
//	}
//	dfs(root, 0, []int{})
//	return
//}

func pathSum(root *TreeNode113, targetSum int) (result [][]int) {
	path := make([]int, 0)
	var dfs func(*TreeNode113, int, []int)
	dfs = func(node *TreeNode113, s int, p []int) {
		if node != nil {
			p = append(p, node.Val)
			s -= node.Val
			if node.Left == nil && node.Right == nil {
				if s == 0 {
					result = append(result, append([]int(nil), p...))
				}
			} else {
				dfs(node.Left, s, p)
				dfs(node.Right, s, p)
				p = p[:len(p)-1]
			}
		}
	}
	dfs(root, targetSum, path)
	return
}

func main() {
	seven := &TreeNode113{7, nil, nil}
	two := &TreeNode113{2, nil, nil}
	five := &TreeNode113{5, nil, nil}
	one := &TreeNode113{1, nil, nil}
	eleven := &TreeNode113{11, seven, two}
	thirteen := &TreeNode113{13, nil, nil}
	four := &TreeNode113{4, five, one}
	four2 := &TreeNode113{4, eleven, nil}
	eight := &TreeNode113{8, thirteen, four}
	five2 := &TreeNode113{5, four2, eight}
	fmt.Println(pathSum(five2, 22))
}
