package main

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 给定的5张牌是否是顺子
 * @param nums int整型一维数组 扑克牌对应的数字集合
 * @return bool布尔型
 */
func isStraight(nums []int) bool {
	sort.Ints(nums)
	zero := 0
	g := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			zero++
		} else {
			diff := nums[i+1] - nums[i]
			if diff == 0 {
				return false
			}
			g += diff - 1
		}
	}
	if zero == 1 {
		if nums[1] == 1 && nums[2] == 11 && nums[3] == 12 && nums[4] == 13 {
			return true
		}
		if nums[1] == 1 && nums[2] == 10 && nums[3] == 12 && nums[4] == 13 {
			return true
		}
		if nums[1] == 1 && nums[2] == 10 && nums[3] == 11 && nums[4] == 13 {
			return true
		}
		if nums[1] == 1 && nums[2] == 10 && nums[3] == 11 && nums[4] == 12 {
			return true
		}

		if nums[1] == 1 && nums[2] == 2 && nums[3] == 11 && nums[4] == 12 {
			return true // 13
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 11 && nums[4] == 13 {
			return true // 12
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 12 && nums[4] == 13 {
			return true // 11
		}
		if nums[1] == 2 && nums[2] == 11 && nums[3] == 12 && nums[4] == 13 {
			return true // 1
		}
		if nums[1] == 1 && nums[2] == 11 && nums[3] == 12 && nums[4] == 13 {
			return true // 2
		}

		if nums[1] == 1 && nums[2] == 2 && nums[3] == 3 && nums[4] == 13 {
			return true // 12
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 3 && nums[4] == 12 {
			return true // 13
		}
		if nums[1] == 2 && nums[2] == 3 && nums[3] == 12 && nums[4] == 13 {
			return true // 1
		}
		if nums[1] == 1 && nums[2] == 3 && nums[3] == 12 && nums[4] == 13 {
			return true // 2
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 12 && nums[4] == 13 {
			return true // 3
		}

		if nums[1] == 1 && nums[2] == 2 && nums[3] == 3 && nums[4] == 4 {
			return true // 13
		}
		if nums[1] == 2 && nums[2] == 3 && nums[3] == 4 && nums[4] == 13 {
			return true // 1
		}
		if nums[1] == 1 && nums[2] == 3 && nums[3] == 4 && nums[4] == 13 {
			return true // 2
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 4 && nums[4] == 13 {
			return true // 3
		}
		if nums[1] == 1 && nums[2] == 2 && nums[3] == 3 && nums[4] == 13 {
			return true // 4
		}
	}
	if zero == 2 {
		if (nums[2] == 1 || nums[2] == 2 || nums[2] == 3) && nums[3] == 12 && nums[4] == 13 {
			return true // 4
		}
		if nums[2] == 1 && nums[3] == 10 && nums[4] == 11 {
			return true // 4
		}
	}
	if zero == 0 {
		return g == 0
	}
	return g <= zero
}

func main() {
	fmt.Println(isStraight([]int{0, 0, 1, 10, 11}))
}
