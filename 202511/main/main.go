package main

import (
	"container/list"
	"fmt"
)

//请用 Go 语言 实现一个支持以下操作的 LRU 缓存（Least Recently Used Cache）：
//Get(key string) (value string, ok bool)
//返回 key 对应的值，并将该 key 标记为最近使用。
//Put(key string, value string)
//插入或更新 key，如果超出容量，则淘汰最久未使用的 key。
//要求：
//支持固定容量；
//最近访问的元素优先保留；
//实现需具备 O(1) 的时间复杂度；

func NewLru(maxCap int) *LruStruct {
	return &LruStruct{
		maxCap: maxCap,
		cache:  make(map[string]*list.Element, maxCap),
		list:   list.New(),
	}
}

type LruStruct struct {
	maxCap int
	cache  map[string]*list.Element
	list   *list.List
	// Sync.Map
}

type kv struct {
	key string
	val string
}

func (l *LruStruct) Get(key string) (string, bool) {
	res, ok := l.cache[key]
	if !ok {
		return "", false
	}
	l.list.MoveToFront(res)

	return res.Value.(*kv).val, true
}

func (l *LruStruct) Put(key, val string) {
	res, ok := l.cache[key]
	if ok {
		res.Value.(*kv).val = val
		l.list.MoveToFront(res)
		return
	}

	// Lock
	newRes := l.list.PushFront(&kv{key, val})
	l.cache[key] = newRes


	if l.list.Len() > l.maxCap {
		tail := l.list.Back()
		if tail != nil {
			l.list.Remove(tail)
			delete(l.cache, tail.Value.(*kv).key)
		}
	}
}

func main() {
	l := NewLru(3)
	l.Put("k1", "v1")
	l.Put("k2", "v2")
	l.Put("k3", "v3")
	l.Put("k4", "v4")
	fmt.Println(l.Get("k1"))
}
