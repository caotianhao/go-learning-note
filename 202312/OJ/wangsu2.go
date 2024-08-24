package main

import (
	"fmt"
	"sort"
)

type treenode struct {
	val   int
	left  *treenode
	right *treenode
}

func getMax(a ...int) int {
	sort.Ints(a)
	return a[len(a)-1]
}

func main() {
	numOfTreeNode := 0
	_, _ = fmt.Scanf("%d", &numOfTreeNode)
	values, parents := make([]int, 0), make([]int, 0)
	d := 0
	for i := 0; i < numOfTreeNode; i++ {
		_, _ = fmt.Scanf("%d", &d)
		values = append(values, d)
	}
	for i := 0; i < numOfTreeNode; i++ {
		_, _ = fmt.Scanf("%d", &d)
		parents = append(parents, d)
	}

	child := make([][2]int, numOfTreeNode)
	for i := 0; i < numOfTreeNode; i++ {
		child[i][0] = -1
		child[i][1] = -1
		if parents[i] > 0 {
			childNode := &child[parents[i]-1]
			if childNode[0] == -1 {
				childNode[0] = i
			} else {
				childNode[1] = i
			}
		}
	}
	ans := values[0]
	dp := make([]int, numOfTreeNode)
	for i := numOfTreeNode - 1; i >= 0; i-- {
		left, right := 0, 0
		if child[i][0] != -1 {
			left = dp[child[i][0]]
		}
		if child[i][1] != -1 {
			right = dp[child[i][1]]
		}
		dp[i] = getMax(getMax(left, right), 0) + values[i]
		ans = getMax(getMax(ans, values[i], values[i]+left, values[i]+right), values[i]+left+right)
	}
	fmt.Println(ans)
}
