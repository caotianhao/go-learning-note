package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param arr int整型二维数组
 * @param n int整型
 * @param m int整型
 * @return int整型一维数组
 */
func println1(arr [][]int, n int, m int) (res []int) {
	for i := 0; i < n; i++ {
		if i&1 == m-m {
			res = append(res, arr[i]...)
		} else {
			res = append(res, reverse1(arr[i])...)
		}
	}
	return
}

func reverse1(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(println1([][]int{{1, 2, 3}}, 3, 6))
}
