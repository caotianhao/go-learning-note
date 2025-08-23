package main

import "fmt"

//func prevPermOpt1(arr []int) []int {
//	l := len(arr)
//	for i := l - 1; i >= 0; i-- {
//		for j := i - 1; j >= 0; j-- {
//			if arr[j] > arr[i] {
//				arr[j], arr[i] = arr[i], arr[j]
//				return arr
//			}
//		}
//	}
//	return arr
//}

func prevPermOpt1(arr []int) []int {
	// l 数组的长度
	// ind 记录从 ind 开始，右面数组为有序的第一个下标
	l, ind, tmp := len(arr), -1, -1
	//fmt.Println("l", l)
	// 如果包含 i 的右面 是有序的
	for i := 0; i < l; i++ {
		if isOrdered(arr[i:l]) {
			//fmt.Println("a", i)
			ind = i
			break
		}
	}
	//fmt.Println("ind", ind)
	// 如果 ind == 0， 那么说明数组本身有序，直接返回
	if ind == 0 {
		return arr
	}
	// 否则从数组末尾倒序遍历，找到严格小于 ind-1 下标的数，先记录
	for i := l - 1; i >= ind; i-- {
		if arr[i] < arr[ind-1] {
			tmp = arr[i]
			break
		}
	}
	//fmt.Println("tmp", tmp)
	// 从 ind 开始，找到等于 tmp 的第一个数字，交换
	// 如果不是这样的话，[3,1,1,3] 就会变成 [1,1,3,3]（下标0和2交换）
	for i := ind; i < l; i++ {
		if arr[i] == tmp {
			arr[ind-1], arr[i] = arr[i], arr[ind-1]
			break
		}
	}
	return arr
}

func isOrdered(arr []int) bool {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
	//fmt.Println(isOrdered([]int{2, 1}))
}
