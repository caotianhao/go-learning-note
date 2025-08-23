package main

import (
	"fmt"
)

type TreeNode654 struct {
	Val   int
	Left  *TreeNode654
	Right *TreeNode654
}

func constructMaximumBinaryTree(nums []int) *TreeNode654 {
	maxB, ind := findMaxAndIndex(nums)
	if ind != -1 {
		root := &TreeNode654{maxB, nil, nil}
		root.Left = constructMaximumBinaryTree(nums[:ind])
		root.Right = constructMaximumBinaryTree(nums[ind+1:])
		return root
	}
	return nil
}

func findMaxAndIndex(nums []int) (int, int) {
	maxI, ind := -1, -1
	for i, v := range nums {
		if v > maxI {
			maxI = v
			ind = i
		}
	}
	return maxI, ind
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	//fmt.Println(findMaxAndIndex([]int{3, 2, 1, 6, 0, 5}))
}
