package main

import "fmt"

// 建图，太傻的做法
func findJudge(n int, trust [][]int) int {
	judge := make([][]int, n)
	for i := 0; i < n; i++ {
		judge[i] = make([]int, n)
	}
	for _, v := range trust {
		judge[v[0]-1][v[1]-1] = 1
	}
	for i := 1; i <= n; i++ {
		f1, f2 := true, true
		for j := 0; j < n; j++ {
			if j != i-1 && judge[j][i-1] != 1 {
				f1 = false
			}
		}
		for j := 0; j < n; j++ {
			if judge[i-1][j] != 0 {
				f2 = false
			}
		}
		if f1 && f2 {
			return i
		}
	}
	return -1
}

// 直接用入度，出度
func findJudge1(n int, trust [][]int) int {
	gIn, gOut := make([]int, n+1), make([]int, n+1)
	for _, v := range trust {
		gIn[v[1]]++
		gOut[v[0]]++
	}
	for i := 1; i <= n; i++ {
		if gIn[i] == n-1 && gOut[i] == 0 {
			return i
		}
	}
	return -1
}

func findJudge2(n int, trust [][]int) int {
	// [out, in]
	find := make([][2]int, n)
	for _, v := range trust {
		find[v[0]-1][0]++
		find[v[1]-1][1]++
	}
	for i, v := range find {
		if v == [2]int{0, n - 1} {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge1(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge2(3, [][]int{{1, 3}, {2, 3}}))
}
