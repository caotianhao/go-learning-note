package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)
	for i, cur := range nums[:n-2] {
		if i > 0 && cur == nums[i-1] {
			continue
		}
		if cur+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if cur+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		nxt, tmp := i+1, n-1
		for nxt < tmp {
			sum := cur + nums[nxt] + nums[tmp]
			if sum > 0 {
				tmp--
			} else if sum < 0 {
				nxt++
			} else {
				ans = append(ans, []int{cur, nums[nxt], nums[tmp]})
				for nxt++; nxt < tmp && nums[nxt] == nums[nxt-1]; nxt++ {
				}
				for tmp--; tmp > nxt && nums[tmp] == nums[tmp+1]; tmp-- {
				}
			}
		}
	}
	return
}

func main() {
	var s string
	nums := make([]int, 0)
	_, _ = fmt.Scanf("%s", &s)
	data := strings.Split(s, ",")
	for _, d := range data {
		t, _ := strconv.Atoi(d)
		nums = append(nums, t)
	}
	ans := threeSum(nums)
	fmt.Printf("[")
	for s := 0; s < len(ans)-1; s++ {
		fmt.Printf("[")
		for i := 0; i < len(ans[s])-1; i++ {
			fmt.Printf("%d, ", ans[s][i])
		}
		fmt.Printf("%d", ans[s][len(ans[s])-1])
		fmt.Printf("], ")
	}
	fmt.Printf("[")
	for i := 0; i < len(ans[len(ans)-1])-1; i++ {
		fmt.Printf("%d, ", ans[len(ans)-1][i])
	}
	fmt.Printf("%d", ans[len(ans)-1][len(ans[len(ans)-1])-1])
	fmt.Printf("]]")
}
