package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/Arafatk/DataViz/trees/redblacktree"
)

// 给你一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。
// 方法 1
func thirdMax1(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	l := len(nums)
	for i, diff := 1, 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}

// 方法 2：有序集合，红黑树
func thirdMax2(nums []int) int {
	// 创建一棵红黑树
	// 为啥用红黑树....因为 go 没有像 C++ 那样的 set
	// 因为 C++ 的 set 也是红黑树实现的
	rbt := redblacktree.NewWithIntComparator()
	// 往里扔数据
	for _, v := range nums {
		rbt.Put(v, nil)
		// 超过三个把当前最小的拿出去
		if rbt.Size() > 3 {
			rbt.Remove(rbt.Left().Key)
		}
	}
	// 如果有 3 个元素那就返回最小的，也就是第三大的
	if rbt.Size() == 3 {
		return rbt.Left().Key.(int)
	}
	// 没 3 个元素的话，返回最大的
	return rbt.Right().Key.(int)
}

// 方法 3：三个变量分别维护最大数，第二最大数，第三最大数
func thirdMax3(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max1, max2, max3 = nums[i], max1, max2
		} else if nums[i] < max1 && nums[i] > max2 {
			max2, max3 = nums[i], max2
		} else if nums[i] < max2 && nums[i] > max3 {
			max3 = nums[i]
		}
	}
	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}

func main() {
	arr := []int{2, 3, 4, 1}
	fmt.Println(thirdMax1(arr))
	fmt.Println(thirdMax2(arr))
	fmt.Println(thirdMax3(arr))
}
