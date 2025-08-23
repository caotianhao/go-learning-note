package main

import "fmt"

type Node590 struct {
	Val      int
	Children []*Node590
}

func postorder590(root *Node590) (res []int) {
	var truePostorder func(node *Node590)
	truePostorder = func(node *Node590) {
		if node != nil {
			for _, v := range node.Children {
				truePostorder(v)
			}
			res = append(res, node.Val)
		}
	}
	truePostorder(root)
	return
}

func main() {
	six := Node590{6, []*Node590{}}
	five := Node590{5, []*Node590{}}
	three := Node590{3, []*Node590{&five, &six}}
	two := Node590{2, []*Node590{}}
	four := Node590{4, []*Node590{}}
	one := Node590{1, []*Node590{&three, &two, &four}}
	fmt.Println(postorder590(&one))
}
