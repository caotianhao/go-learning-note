package main

import "fmt"

type Node589 struct {
	Val      int
	Children []*Node589
}

func preorder589(root *Node589) (res []int) {
	var truePreorder func(node *Node589)
	truePreorder = func(node *Node589) {
		if node != nil {
			res = append(res, node.Val)
			for _, v := range node.Children {
				truePreorder(v)
			}
		}
	}
	truePreorder(root)
	return
}

func main() {
	six := Node589{6, []*Node589{}}
	five := Node589{5, []*Node589{}}
	three := Node589{3, []*Node589{&five, &six}}
	two := Node589{2, []*Node589{}}
	four := Node589{4, []*Node589{}}
	one := Node589{1, []*Node589{&three, &two, &four}}
	fmt.Println(preorder589(&one))
}
