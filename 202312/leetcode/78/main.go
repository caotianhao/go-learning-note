package main

import "fmt"

//func subsets(nums []int) (res [][]int) {
//	l := len(nums)
//	for i := 0; i < 1<<l; i++ {
//		tmp := make([]int, 0)
//		for j := 0; j < i; j++ {
//			if i>>j&1 == 1 {
//				tmp = append(tmp, nums[l-j-1])
//			}
//		}
//		res = append(res, tmp)
//	}
//	return
//}

func subsets(nums []int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		if len(nums) == i {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
		dfs(i + 1)
	}
	dfs(0)
	return
}

func main() {
	fmt.Println(subsets([]int{122, 333}))
}
