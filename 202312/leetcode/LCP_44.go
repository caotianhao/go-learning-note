package main

import "fmt"

type TreeNodeLCP44 struct {
	Val   int
	Left  *TreeNodeLCP44
	Right *TreeNodeLCP44
}

func numColor(root *TreeNodeLCP44) (cnt int) {
	return len(inOrder44(root))
}

func inOrder44(root *TreeNodeLCP44) map[int]struct{} {
	res := map[int]struct{}{}
	var trueInorderTraversal func(node *TreeNodeLCP44)
	trueInorderTraversal = func(node *TreeNodeLCP44) {
		if node != nil {
			trueInorderTraversal(node.Left)
			res[node.Val] = struct{}{}
			trueInorderTraversal(node.Right)
		}
	}
	trueInorderTraversal(root)
	return res
}

func main() {
	l := TreeNodeLCP44{3, nil, nil}
	r := TreeNodeLCP44{3, nil, nil}
	root := TreeNodeLCP44{3, &l, &r}
	fmt.Println(numColor(&root))
}
