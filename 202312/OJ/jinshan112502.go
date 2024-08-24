package main

import (
	"container/heap"
	"container/list"
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
)

type j05 struct {
	sort.IntSlice
}

func (j *j05) Push(v any) {
	j.IntSlice = append(j.IntSlice, v.(int))
}

func (j *j05) Pop() any {
	o := j.IntSlice[len(j.IntSlice)-1]
	j.IntSlice = j.IntSlice[:len(j.IntSlice)-1]
	return o
}

func j03(a int) int {
	if a > 5 {
		return 9
	}
	return 3
}

func j04(stack list.List) int {
	return stack.Len()
}

func main() {
	a := http.StatusOK
	b := math.MaxInt64
	if j03(a) == 14 {
		log.Printf("aaa")
	} else {
		fmt.Println("sd" + fmt.Sprintf("%d", b))
	}
	c := list.List{}
	fmt.Println(j04(c))
	j := &j05{}
	heap.Init(j)
	q := context.Background()
	fmt.Println(q)
}
