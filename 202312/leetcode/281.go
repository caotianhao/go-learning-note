package main

import "fmt"

type ZigzagIterator struct {
	a1, a2 []int
	l      int
	cnt    int
}

func Constructor281(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		a1:  v1,
		a2:  v2,
		l:   len(v1) + len(v2),
		cnt: 0,
	}
}

func (zi *ZigzagIterator) next() int {
	l1, l2 := len(zi.a1), len(zi.a2)
	var r int
	if zi.cnt < min(l1, l2)*2 {
		if zi.cnt%2 == 0 {
			r = zi.a1[zi.cnt/2]
		} else {
			r = zi.a2[zi.cnt/2]
		}
		zi.cnt++
	} else {
		if l1 > l2 {
			r = zi.a1[zi.cnt-l2]
			zi.cnt++
		} else if l1 < l2 {
			r = zi.a2[zi.cnt-l1]
			zi.cnt++
		}
	}
	return r
}

func (zi *ZigzagIterator) hasNext() bool {
	return zi.cnt < zi.l
}

func main() {
	c := Constructor281([]int{1, 2, 7, 8, 9}, []int{3, 4, 5, 6})
	for c.hasNext() {
		fmt.Printf("%d\t", c.next())
	}
}
