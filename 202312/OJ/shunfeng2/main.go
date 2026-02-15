package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	var n, m, ability int64
	_, _ = fmt.Scanf("%d %d", &n, &m)
	ab := make([]int64, 0)
	for i := int64(0); i < n; i++ {
		_, _ = fmt.Scanf("%d", &ability)
		ab = append(ab, ability)
	}
	sort.Slice(ab, func(i, j int) bool {
		return ab[i] > ab[j]
	})
	var queue list.List
	for i := 0; i < len(ab); i++ {
		queue.PushBack(ab[i])
	}
	group := 0
	for i := 0; i < len(ab); i++ {
		if ab[i] >= m {
			group++
			queue.Remove(queue.Front())
		}
	}
	for queue.Len() != 0 {
		t := queue.Front().Value.(int64)
		var id int64
		for flag1 := int64(2); ; flag1++ {
			if flag1*t >= m {
				id = flag1
				break
			}
		}
		if int64(queue.Len()) >= id {
			for move := int64(0); move < id; move++ {
				queue.Remove(queue.Back())
			}
			group++
		}
	}
	fmt.Println(group)
}
