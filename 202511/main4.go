package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func preOrder(root *Node) (res []int) {
	if root == nil {
		return []int{}
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

func main() {

	n2 := Node{2, nil, nil}
	n3 := Node{3, nil, nil}
	n1 := Node{1, &n2, &n3}
	fmt.Println(preOrder(&n1))

	//    1
	// 	  2 3
	// 4 5 6 7
	n2.Left = &Node{4, nil, nil}
	n2.Right = &Node{5, nil, nil}
	n3.Left = &Node{6, nil, nil}
	n3.Right = &Node{7, nil, nil}
	fmt.Println(preOrder(&n1))
}
