package main

import "fmt"

//哈希表法
//func containsNearbyDuplicate(nums []int, k int) bool {
//	hashMap := map[int]int{}
//	//利用哈希表存储元素，key 为值，value 为对应下标
//	for i, v := range nums {
//		//如果元素已经存在于哈希表中，也就是 ok 为 true
//		//且元素原来的位置 pos 和新位置 i （一定比原位置大，因为是从头遍历）的差 <= k，则返回 true
//		if pos, ok := hashMap[v]; ok && (i-pos <= k) {
//			fmt.Println(i, pos)
//			return true
//		}
//		//正常往里存
//		hashMap[v] = i
//	}
//	return false
//}

//暴力解法
//func containsNearbyDuplicate(nums []int, k int) bool {
//	l := len(nums)
//	for i := 0; i < l; i++ {
//		//j = i+1，j 一定大于 i，所以原题的意思可以变成 j-i <= k，也就是 j <= k+i
//		//同时保证基本的数组不越界，加上 j < l
//		for j := i + 1; j <= k+i && j < l; j++ {
//			if nums[i] == nums[j] {
//				return true
//			}
//		}
//	}
//	return false
//}

//滑动窗口法
func containsNearbyDuplicate(nums []int, k int) bool {
	//用 set 存储，也是 key 对应元素，value 对应下标
	set := map[int]int{}
	for i, num := range nums {
		//滑动窗口显然大小应该为 k+1，当窗口满了的时候，先删除最左面元素
		if i > k {
			delete(set, nums[i-k-1])
		}
		//如果新元素已经在删除过旧元素之后的 set 中，那么一定符合题意
		if _, ok := set[num]; ok {
			return true
		}
		//否则继续追加
		set[num] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
