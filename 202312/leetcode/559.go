package main

import "fmt"

type Node559 struct {
	Val      int
	Children []*Node559
}

func maxDepth559(root *Node559) int {
	if root == nil {
		return 0
	}
	maxD := 0
	for _, v := range root.Children {
		tmp := maxDepth559(v)
		if tmp > maxD {
			maxD = tmp
		}
	}
	return maxD + 1
}

func main() {
	six := Node559{6, []*Node559{}}
	five := Node559{5, []*Node559{}}
	three := Node559{3, []*Node559{&five, &six}}
	two := Node559{2, []*Node559{}}
	four := Node559{4, []*Node559{}}
	one := Node559{1, []*Node559{&three, &two, &four}}
	fmt.Println(maxDepth559(&one))
}
