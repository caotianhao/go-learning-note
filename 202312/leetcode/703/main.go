package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// KthLargest 堆正式入门题
// 正常来讲一个堆需要实现 Swap, Len, Less, Pop, Push 方法
// 但 sort.IntSlice 自带前三个，所以实现堆可以直接使用
// 这样可以只实现 Pop 和 Push 方法
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor703(k int, nums []int) KthLargest {
	// 直接使用 Add 方法增加
	// 如果直接把 nums 加进来，还要看 nums 的长度和 k 的关系
	// 这里很巧妙
	res := KthLargest{k: k}
	for _, v := range nums {
		res.Add(v)
	}
	return res
}

func (kl *KthLargest) Add(val int) int {
	// 默认是一个小根堆，其实这里和红黑树的不同之处就在于
	// 红黑树是 set，而本题可以有重复元素
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	// return 的也要注意，堆说白了只有堆顶有用
	return kl.IntSlice[0]
}

// Push 自己实现的
// 注意必须是接口类型，需要什么类型自己断言
func (kl *KthLargest) Push(v any) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

// Pop 返回值类型也必须是接口
// 若后续需要断言可继续操作，但返回值必须是接口
func (kl *KthLargest) Pop() any {
	res := kl.IntSlice[len(kl.IntSlice)-1]
	kl.IntSlice = kl.IntSlice[:len(kl.IntSlice)-1]
	return res
}

func main() {
	kk := Constructor703(3, []int{4, 5, 8, 2})
	fmt.Println(kk.Add(3))
	fmt.Println(kk.Add(5))
	fmt.Println(kk.Add(10))
	fmt.Println(kk.Add(9))
	fmt.Println(kk.Add(4))
	fmt.Println("----------------------------------------------")
	k := Constructor703(1, []int{})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
}
