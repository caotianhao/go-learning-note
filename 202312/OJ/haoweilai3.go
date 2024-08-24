package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type listNode struct {
	value int
	next  *listNode
}

type nodeHeap []*listNode

func (h *nodeHeap) Len() int {
	return len(*h)
}

func (h *nodeHeap) Less(i, j int) bool {
	return (*h)[i].value < (*h)[j].value
}

func (h *nodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *nodeHeap) Push(v interface{}) {
	*h = append(*h, v.(*listNode))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func mergeKAseLists(heads []*listNode) *listNode {
	h := nodeHeap{}
	for _, head := range heads {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)
	dummyNode := &listNode{}
	cur := dummyNode
	for len(h) > 0 {
		node := heap.Pop(&h).(*listNode)
		if node.next != nil {
			heap.Push(&h, node.next)
		}
		cur.next = node
		cur = cur.next
	}
	return dummyNode.next
}

func main() {
	numberOfLists := 0
	_, _ = fmt.Scanf("%d", &numberOfLists)
	lists := make([]*listNode, 0)
	var s string
	for i := 0; i < numberOfLists; i++ {
		head := &listNode{}
		cur := head
		_, _ = fmt.Scanf("%s", &s)
		data := strings.Split(s, ",")
		for _, v := range data {
			val, _ := strconv.Atoi(v)
			t := &listNode{value: val}
			cur.next = t
			cur = cur.next
		}
		lists = append(lists, head.next)
	}
	ans := mergeKAseLists(lists)
	dummyHead, l := ans, 0
	for dummyHead != nil {
		l++
		dummyHead = dummyHead.next
	}
	for i := 0; i < l-1; i++ {
		fmt.Printf("%d,", ans.value)
		ans = ans.next
	}
	fmt.Printf("%d", ans.value)
}
