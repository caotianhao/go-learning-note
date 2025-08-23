package main

import "fmt"

func permute2083(nums []int) (ans [][]int) {
	l := len(nums)
	used, path := make([]bool, l), make([]int, l)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j, is := range used {
			if !is {
				path[i] = nums[j]
				used[j] = true
				dfs(i + 1)
				used[j] = false
			}
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Println(permute2083([]int{2, 5, 6}))
}
