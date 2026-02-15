package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNodeYaXin struct {
	Val   int
	Left  *TreeNodeYaXin
	Right *TreeNodeYaXin
}

func buildTree(arr []int, index, n int) *TreeNodeYaXin {
	if n == 0 || index >= n {
		return nil
	}
	return &TreeNodeYaXin{
		Val:   arr[index],
		Left:  buildTree(arr, index*2+1, n),
		Right: buildTree(arr, index*2+2, n),
	}
}

func inOrder(root *TreeNodeYaXin) (res []int) {
	var dfs func(*TreeNodeYaXin)
	dfs = func(node *TreeNodeYaXin) {
		if node != nil {
			dfs(node.Left)
			res = append(res, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	arr := make([]int, 0)
	data := strings.Split(s, ",")
	for _, d := range data {
		t, _ := strconv.Atoi(d)
		arr = append(arr, t)
	}
	root := buildTree(arr, 0, len(arr))
	result := inOrder(root)
	for _, v := range result {
		fmt.Printf("%d,", v)
	}
}
