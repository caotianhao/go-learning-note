package main

import "fmt"

//func findKthPositive(arr []int, k int) int {
//	hashMap := map[int]struct{}{}
//	for _, v := range arr {
//		hashMap[v] = struct{}{}
//	}
//	for i := 1; i <= 2000; i++ {
//		if _, ok := hashMap[i]; !ok {
//			k--
//		}
//		if k == 0 {
//			return i
//		}
//	}
//	return -1
//}

func findKthPositive(arr []int, k int) int {
	l := len(arr)
	// 为了防止特判，在数组后面加一个贼大的数，由题意这个数是2001
	arr = append(arr, 2001)
	// 二分查找的左中右
	left, right, mid := 0, l, 0
	for left <= right {
		mid = (left + right) / 2
		// 当前元素 - 下标 - 1
		// 就是截止到当前元素为止，缺失的正整数个数
		// 如果这个结果比k大了，那么就说明找到答案这里了
		if arr[mid]-mid-1 >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 其实应该是 return k - (arr[left] - (left - 1) - 1) + arr[left]
	// 结果化简为 k + left
	return k + left
}

func main() {
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 5))
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}
