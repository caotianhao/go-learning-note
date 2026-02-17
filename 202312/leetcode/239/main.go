package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 堆正式入门
// 定义全局整型数组
// 为了记录下标和排序下标
// 如果不这样，而是直接把数组元素传入堆中，那么假设下一步窗口移动的时候
// 最大值被移出去，在堆中无法操作
var publicSlice []int

// 堆
type hp struct {
	sort.IntSlice
}

// Less 这里在 hp 前面加 * 和不加只要不是引入接口就没有区别
// 而不加 * 会在 GoLand 中报错
func (h *hp) Less(i, j int) bool {
	//return a[h.IntSlice[i]] > a[h.IntSlice[j]]
	return publicSlice[h.IntSlice[i]] > publicSlice[h.IntSlice[j]]
}

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() any {
	ans := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return ans
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	// 这一步必须要有，这样才能对 nums 排序
	publicSlice = nums
	// 初始化建堆
	myHeap := &hp{}
	// 堆中存储下标，一定要存储下标
	for i := 0; i < k; i++ {
		myHeap.IntSlice = append(myHeap.IntSlice, i)
		//heap.Push(myHeap, i)
	}
	// 初始化
	heap.Init(myHeap)
	//fmt.Println(myHeap)
	// 由于是按下标所代表的数值排序，因此堆顶就是当前最大的元素的下标
	ans = append(ans, nums[myHeap.IntSlice[0]])
	// 后面的元素进堆
	for i := k; i < len(nums); i++ {
		heap.Push(myHeap, i)
		//fmt.Println("before", myHeap)
		// 如果堆顶元素对应的下标不在窗口该有的下标范围内，就弹出
		for myHeap.IntSlice[0] <= i-k {
			heap.Pop(myHeap)
		}
		//fmt.Println("after", myHeap)
		// 然后继续在答案数组中加入堆顶元素下标对应的值
		ans = append(ans, nums[myHeap.IntSlice[0]])
	}
	return
}

func main() {
	// 3,3,5,5,5,4
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, -2, 4}, 3))
}
