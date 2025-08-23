package main

import "fmt"

type TreeNode508 struct {
	Val   int
	Left  *TreeNode508
	Right *TreeNode508
}

// 暴力
//func dfs508(root *TreeNode508) (s int) {
//	var cal func(node *TreeNode508)
//	cal = func(node *TreeNode508) {
//		if node != nil {
//			cal(node.Left)
//			s += node.Val
//			cal(node.Right)
//		}
//	}
//	cal(root)
//	return
//}
//
//func findFrequentTreeSum(root *TreeNode508) (r []int) {
//	var getEveryNodeSum func(node *TreeNode508)
//	sumArr := make([]int, 0)
//	getEveryNodeSum = func(node *TreeNode508) {
//		if node != nil {
//			getEveryNodeSum(node.Left)
//			sumArr = append(sumArr, dfs508(node))
//			getEveryNodeSum(node.Right)
//		}
//	}
//	getEveryNodeSum(root)
//	m, most := make(map[int]int), -1
//	for _, v := range sumArr {
//		m[v]++
//	}
//	for _, v := range m {
//		if v > most {
//			most = v
//		}
//	}
//	for i, v := range m {
//		if v == most {
//			r = append(r, i)
//		}
//	}
//	return
//}

func findFrequentTreeSum(root *TreeNode508) (ans []int) {
	var dfs func(node *TreeNode508) int
	cnt, maxCnt := make(map[int]int), -1
	dfs = func(node *TreeNode508) int {
		if node == nil {
			return 0
		}
		s := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[s]++
		if cnt[s] > maxCnt {
			maxCnt = cnt[s]
		}
		return s
	}
	dfs(root)
	for val, count := range cnt {
		if count == maxCnt {
			ans = append(ans, val)
		}
	}
	return
}

func main() {
	n3 := &TreeNode508{-3, nil, nil}
	two := &TreeNode508{2, nil, nil}
	five := &TreeNode508{5, two, n3}
	fmt.Println(findFrequentTreeSum(five))

	n5 := &TreeNode508{-5, nil, nil}
	two1 := &TreeNode508{2, nil, nil}
	five1 := &TreeNode508{5, two1, n5}
	fmt.Println(findFrequentTreeSum(five1))
}
