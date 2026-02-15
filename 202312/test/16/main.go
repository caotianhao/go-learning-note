package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	partSum := 0
	n := len(nums)
	ans := math.MaxInt64
	for ; right < n; right++ {
		partSum += nums[right]
		for partSum >= target {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			partSum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

//func main() {
//	//nums := make([]int, 0)
//	var data []string
//	cnt := 0
//	input := bufio.NewScanner(os.Stdin)
//	for input.Scan() {
//		data = strings.Split(input.Text(), ",")
//		cnt = len(data)
//		//for _, d := range data {
//		//t, _ := strconv.Atoi(d)
//		//cnt++
//		//nums = append(nums, t)
//		//}
//	}
//	nums := make([]int, cnt)
//	for _, d := range data {
//		t, _ := strconv.Atoi(d)
//		nums = append(nums, t)
//	}
//	target := 0
//	_, _ = fmt.Scanf("%d", &target)
//	fmt.Println(minSubArrayLen(target, nums))
//	//fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
//}

func main() {
	nums := make([]int, 0)
	var s string
	_, _ = fmt.Scanf("%s", &s)
	target := 0
	_, _ = fmt.Scanf("%d", &target)
	fmt.Println(s)
	data := strings.Split(s, ",")
	for _, v := range data {
		t, _ := strconv.Atoi(v)
		nums = append(nums, t)
	}
	fmt.Println("g", target, nums)
	fmt.Println(minSubArrayLen(target, nums))
}
