package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func CreateBinaryTree(data int) *BinaryTreeNode {
	return &BinaryTreeNode{data, nil, nil}
}

func (node *BinaryTreeNode) Insert(n *BinaryTreeNode, data int) bool {
	cur := n
	for cur != nil {
		if cur.Data < data {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = CreateBinaryTree(data)
				return true
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = CreateBinaryTree(data)
				fmt.Println(data, "d---")
				return true
			}
		}
	}
	return false
}

func (node *BinaryTreeNode) BreadthFirstSearch() []int {
	if node == nil {
		return nil
	}
	var result []int
	par := node
	cur := []*BinaryTreeNode{par}
	for len(cur) > 0 {
		result = append(result, cur[0].Data)
		if cur[0].Left != nil {
			cur = append(cur, cur[0].Left)
		}
		if cur[0].Right != nil {
			cur = append(cur, cur[0].Right)
		}
		cur = cur[1:]
	}
	return result
}

func (node *BinaryTreeNode) PreOrder(n *BinaryTreeNode) {
	if n != nil {
		fmt.Println(n.Data)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

func (node *BinaryTreeNode) InOrder(n *BinaryTreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		fmt.Println(n.Data)
		node.InOrder(n.Right)
	}
}

func (node *BinaryTreeNode) PostOrder(n *BinaryTreeNode) {
	if n != nil {
		node.InOrder(n.Left)
		node.InOrder(n.Right)
		fmt.Println(n.Data)
	}
}

func (node *BinaryTreeNode) GetHeight(n *BinaryTreeNode) int {
	if n == nil {
		return 0
	}
	l := node.GetHeight(n.Left)
	r := node.GetHeight(n.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func (node *BinaryTreeNode) FindLead(n *BinaryTreeNode) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			fmt.Println(n.Data)
		}
		node.FindLead(n.Left)
		node.FindLead(n.Right)
	}
}

func (node *BinaryTreeNode) FindValueNode(n *BinaryTreeNode, target int) *BinaryTreeNode {
	if n == nil {
		return nil
	} else if n.Data == target {
		return n
	} else {
		cur := node.FindValueNode(n.Left, target)
		if cur != nil {
			return cur
		}
		return node.FindValueNode(n.Right, target)
	}
}

func main() {
	var node *BinaryTreeNode
	node = CreateBinaryTree(10)
	li := []int{2, 1, 3}
	for _, val := range li {
		node.Insert(node, val)
	}
	//ret := node.BreadthFirstSearch()
	//fmt.Println("ret =", ret)
	//node.PreOrder(node)
	//node.InOrder(node)
	//node.PostOrder(node)
	res := node.GetHeight(node)
	fmt.Println("GetHeight(node)", res)
	//node.FindLead(node)
	//ref := node.FindValueNode(node, 17)
	//fmt.Println(ref)
}
