package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//题目要求：
//给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数大于 n/2 的元素
//你可以假设数组是非空的，并且给定的数组总是存在多数元素

// 方法一：哈希表
// 自己做的，时间空间都是 O(n)
func majorityElement1(nums []int) int {
	map169, ind := map[int]int{}, -1
	for _, v := range nums {
		map169[v]++
	}
	for i, v := range map169 {
		if v > len(nums)/2 {
			ind = i
			break
		}
	}
	return ind
}

// 方法二：排序
// 这样下标为 n/2 的元素一定满足条件
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 方法三：随机
// 因为符合条件的数在数组中至少占一半，因此随机选个数，判断它是不是，是就 return，不是就继续
// 最坏的情况下时间复杂度为无穷大，正常情况时间复杂度为 O(n),空间复杂度为 O(1)
func majorityElement3(nums []int) int {
	l, cnt := len(nums), 0
	rand.NewSource(time.Now().UnixNano())
	for {
		n := nums[rand.Int()%l]
		for _, v := range nums {
			if v == n {
				cnt++
			}
			if cnt > l/2 {
				return n
			}
		}
	}
}

//方法四：分治

// 方法五：Boyer-Moore 投票算法
// 维护一个候选众数 candidate 和它出现的次数 count
// 初始时 candidate 可以为任意值，count 为 0
// 遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，如果 count 的值为 0，我们先将 x 的值赋予 candidate
// 随后判断 x
// 如果 x 与 candidate 相等，那么计数器 count 的值增加 1
// 如果 x 与 candidate 不等，那么计数器 count 的值减少 1
// 在遍历完成后，candidate 即为整个数组的众数
func majorityElement5(nums []int) int {
	candidate, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			candidate = v
		}
		if v == candidate {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	slice169 := make([]int, 0)

	//题目中数组大小最大为 50000
	for i := 0; i < 24000; i++ {
		slice169 = append(slice169, rand.Intn(1000))
	}
	myRand := rand.Intn(1000)
	for i := 0; i < 26000; i++ {
		slice169 = append(slice169, myRand)
	}

	time1begin := time.Now()
	a1 := majorityElement1(slice169)
	time1end := time.Since(time1begin)
	fmt.Println("方法1：", a1, " 时间：", time1end)

	time2begin := time.Now()
	a2 := majorityElement2(slice169)
	time2end := time.Since(time2begin)
	fmt.Println("方法2：", a2, " 时间：", time2end)

	time3begin := time.Now()
	a3 := majorityElement3(slice169)
	time3end := time.Since(time3begin)
	fmt.Println("方法3：", a3, " 时间：", time3end)

	time5begin := time.Now()
	a5 := majorityElement5(slice169)
	time5end := time.Since(time5begin)
	fmt.Println("方法5：", a5, " 时间：", time5end)
}
