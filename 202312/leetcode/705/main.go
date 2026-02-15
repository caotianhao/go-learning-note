package main

import "fmt"

type MyHashSet struct {
	key []int
}

func Constructor705() MyHashSet {
	return MyHashSet{key: []int{}}
}

func (mh *MyHashSet) Add(key int) {
	if isIn705(mh.key, key) {
		return
	}
	mh.key = append(mh.key, key)
}

func (mh *MyHashSet) Remove(key int) {
	if isIn705(mh.key, key) {
		mh.key = delete705(mh.key, key)
	}
}

func (mh *MyHashSet) Contains(key int) bool {
	return isIn705(mh.key, key)
}

func isIn705(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func delete705(arr []int, n int) []int {
	slice705 := make([]int, 0)
	for _, v := range arr {
		if v == n {
			continue
		}
		slice705 = append(slice705, v)
	}
	return slice705
}

func main() {
	//["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
	//[[],            [1],   [2],    [1],         [3],      [2],       [2],      [2],      [2]]
	a := Constructor705()
	fmt.Println("Constructor()", a.key)
	a.Add(1)
	fmt.Println("Add(1)       ", a.key)
	a.Add(2)
	fmt.Println("Add(2)       ", a.key)
	a1 := a.Contains(1)
	fmt.Println("Contains(1)  ", a1)
	a2 := a.Contains(3)
	fmt.Println("Contains(3)  ", a2)
	a.Add(2)
	fmt.Println("Add(2)       ", a.key)
	a3 := a.Contains(2)
	fmt.Println("Contains(2)  ", a3)
	a.Remove(2)
	fmt.Println("Remove(2)    ", a.key)
	a4 := a.Contains(2)
	fmt.Println("Contains(2)  ", a4)
	//fmt.Println(delete705([]int{1, 2, 3, 4, 5}, 5))
}
