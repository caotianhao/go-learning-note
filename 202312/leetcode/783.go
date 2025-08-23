package main

import (
	"fmt"
	"sort"
)

type TreeNode783 struct {
	Val   int
	Left  *TreeNode783
	Right *TreeNode783
}

func minDiffInBST(root *TreeNode783) int {
	arr, minK := make([]int, 0), 999999
	var po = func(node *TreeNode783) {}
	po = func(node *TreeNode783) {
		if node != nil {
			po(node.Left)
			arr = append(arr, node.Val)
			po(node.Right)
		}
	}
	po(root)
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		tmp := arr[i+1] - arr[i]
		if tmp < minK {
			minK = tmp
			if minK == 1 {
				return 1
			}
		}
	}
	return minK
}

func main() {
	one := TreeNode783{1, nil, nil}
	three := TreeNode783{3, nil, nil}
	two := TreeNode783{2, &one, &three}
	six := TreeNode783{6, nil, nil}
	four := TreeNode783{4, &two, &six}
	fmt.Println(minDiffInBST(&four))
}
