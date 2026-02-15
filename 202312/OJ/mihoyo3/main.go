package main

import (
	"fmt"
	"math"
)

//type node struct {
//	ch       [2]*node
//	priority int
//	key      int
//	cnt      int
//}
//
//func (o *node) cmp(b int) int {
//	switch {
//	case b < o.key:
//		return 0
//	case b > o.key:
//		return 1
//	default:
//		return -1
//	}
//}
//
//func (o *node) rotate(d int) *node {
//	x := o.ch[d^1]
//	o.ch[d^1] = x.ch[d]
//	x.ch[d] = o
//	return x
//}
//
//type treap struct {
//	root *node
//}
//
//func (t *treap) ins(o *node, key int) *node {
//	if o == nil {
//		return &node{
//			priority: rand.Int(),
//			key:      key,
//			cnt:      1,
//		}
//	}
//	if d := o.cmp(key); d >= 0 {
//		o.ch[d] = t.ins(o.ch[d], key)
//		if o.ch[d].priority > o.priority {
//			o = o.rotate(d ^ 1)
//		}
//	} else {
//		o.cnt++
//	}
//	return o
//}
//
//func (t *treap) del(o *node, key int) *node {
//	if o == nil {
//		return nil
//	}
//	if d := o.cmp(key); d >= 0 {
//		o.ch[d] = t.del(o.ch[d], key)
//	} else {
//		if o.cnt > 1 {
//			o.cnt--
//		} else {
//			if o.ch[1] == nil {
//				return o.ch[0]
//			}
//			if o.ch[0] == nil {
//				return o.ch[1]
//			}
//			d = 0
//			if o.ch[0].priority > o.ch[1].priority {
//				d = 1
//			}
//			o = o.rotate(d)
//			o.ch[d] = t.del(o.ch[d], key)
//		}
//	}
//	return o
//}
//
//func (t *treap) insert(key int) {
//	t.root = t.ins(t.root, key)
//}
//
//func (t *treap) delete(key int) {
//	t.root = t.del(t.root, key)
//}
//
//func (t *treap) min() (min *node) {
//	for o := t.root; o != nil; o = o.ch[0] {
//		min = o
//	}
//	return
//}
//
//func (t *treap) max() (max *node) {
//	for o := t.root; o != nil; o = o.ch[1] {
//		max = o
//	}
//	return
//}
//
//func mihoyo3(nums []int, limit int) (ans int) {
//	t := &treap{}
//	//mod := int(1e9 + 7)
//	left := 0
//	for _, v := range nums {
//		t.insert(v)
//		for t.max().key-t.min().key >= limit {
//			ans += t.max().key - t.min().key
//			//if ans > mod {
//			//	ans %= mod
//			//}
//			t.delete(nums[left])
//			left++
//		}
//	}
//	return
//}

func mihoyo3(nums []int, limit int) (ans int) {
	m := int(1e9 + 7)
	var minQ, maxQ []int
	left := 0
	for _, v := range nums {
		for len(minQ) > 0 {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			ans += maxQ[0] - minQ[0]
			if ans > m {
				ans %= m
			}
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
	}
	return
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	num := 0
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		a = append(a, num)
	}
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < a[i]; j++ {
			b = append(b, i+1)
		}
	}
	fmt.Println(mihoyo3(b, math.MinInt64))
}
